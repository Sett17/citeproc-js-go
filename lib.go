/*
Package citeproc_js_go

Package citeproc_js_go implements the integration of citeproc-js with Go.
Citeproc-js is a JavaScript implementation of the Citation Style Language (CSL)
that is used to format bibliographic references.

The package allows the user to pass in a CSL file and a locales file,
or if not provided, use the default files included in the package.
*/
package citeproc_js_go

import (
	_ "embed"
	"encoding/xml"
	"fmt"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	_ "github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
	"github.com/sett17/citeproc-js-go/csljson"
	"os"
)

//go:embed citeproc.min.js
var citeprocMinJs string

//go:embed ieee.csl
var ieeeCsl string

//go:embed locales-en-US.xml
var enUsLocale string

// Session is a struct that stores information about the Citeproc session.
type Session struct {
	vm         *goja.Runtime
	cslFile    string
	localeFile string
}

// NewSession creates and returns a new Citeproc session.
func NewSession() *Session {
	return &Session{
		vm: goja.New(),
	}
}

// SetCslContent sets the content of the CSL file for the session.
func (s *Session) SetCslContent(cslContent string) {
	s.cslFile = cslContent
}

// SetCslFile sets the content of the CSL file for the session.
func (s *Session) SetCslFile(cslFilePath string) error {
	buf, err := os.ReadFile(cslFilePath)
	if err != nil {
		return err
	}
	s.cslFile = string(buf)
	return nil
}

// SetLocaleContent sets the content of the locale file for the session.
func (s *Session) SetLocaleContent(localeContent string) {
	s.localeFile = localeContent
}

// SetLocaleFile sets the content of the locale file for the session.
func (s *Session) SetLocaleFile(localeFilePath string) error {
	buf, err := os.ReadFile(localeFilePath)
	if err != nil {
		return err
	}
	s.localeFile = string(buf)
	return nil
}

// Init initializes the Citeproc session by loading required files and setting up the runtime.
func (s *Session) Init() (err error) {
	s.vm.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))
	new(require.Registry).Enable(s.vm)
	console.Enable(s.vm)
	// If cslFile is not set, use the default IEEE CSL file.
	if s.cslFile == "" {
		s.cslFile = ieeeCsl
	}
	err = s.vm.Set("cslFile", s.cslFile)
	if err != nil {
		return err
	}

	// If localeFile is not set, use the default en-US locale file.
	if s.localeFile == "" {
		s.localeFile = enUsLocale
	}
	err = s.vm.Set("localeFile", s.localeFile)
	if err != nil {
		return err
	}

	// Load citeproc-js into the runtime.
	_, err = s.vm.RunString(citeprocMinJs)
	if err != nil {
		return err
	}

	// Set up the citeprocSys object and create a new CSL engine.
	_, err = s.vm.RunString(`
let citations = {};
//let itemIDs = [];

function addCitation(id, citation) {
	citations[id] = citation;
	//itemIDs.push(id);
}

citeprocSys = {
    retrieveLocale: function (lang) {
		return localeFile;
    },
    retrieveItem: function(id){
		//if (id in citations) {
		return citations[id];
		//}
		//return {};
    }
};

let engine = new CSL.Engine(citeprocSys, cslFile);

function pcc(citationData, pre) {
	return engine.processCitationCluster(citationData, pre, []);
}
`)
	if err != nil {
		return err
	}

	return nil
}

func (s *Session) AddCitation(items ...csljson.Item) error {
	addCitation, ok := goja.AssertFunction(s.vm.Get("addCitation"))
	if !ok {
		return fmt.Errorf("addCitation is not a function")
	}

	for i := range items {
		_, err := addCitation(goja.Undefined(), s.vm.ToValue(items[i].CitationKey), s.vm.ToValue(items[i]))
		if err != nil {
			return err
		}
	}

	return nil
}

var citationsPre = make([][]string, 0)

// ProcessCitationCluster processes a citation cluster and returns the resulting string that should be placed in the text at the place of the citation.
func (s *Session) ProcessCitationCluster(items ...csljson.Item) (string, error) {

	citeItems := make([]csljson.CiteItem, len(items))
	for i := range items {
		citeItems[i] = csljson.CiteItemFromItem(items[i])
	}

	citation := csljson.Citation{
		CitationItems: citeItems,
		Properties: struct {
			NoteIndex int `json:"noteIndex"`
		}{
			NoteIndex: 0,
		},
		//CitationID:  fmt.Sprintf("%X", rand.Int31()),
		CitationID:  nil,
		SortedItems: nil,
	}
	//citationsPre = append(citationsPre, []string{citation.CitationID.(string), "0"})
	pcc, ok := goja.AssertFunction(s.vm.Get("pcc"))
	if !ok {
		return "", fmt.Errorf("pcc is not a function")
	}
	result, err := pcc(goja.Undefined(), s.vm.ToValue(citation), s.vm.ToValue(citationsPre))
	if err != nil {
		return "", err
	}
	resExport := result.Export()

	if resExport.([]interface{})[0].(map[string]interface{})["bibchange"].(bool) {
		citationsPre = append(citationsPre, []string{resExport.([]interface{})[1].([]interface{})[0].([]interface{})[2].(string), "0"})
	}
	return resExport.([]interface{})[1].([]interface{})[0].([]interface{})[1].(string), nil
}

// MakeBibliography processes the bibliography items stored in the current session and returns a slice of strings containing the formatted bibliography entries.
func (s *Session) MakeBibliography() ([]string, error) {
	ret := make([]string, 0)
	result, err := s.vm.RunString(`engine.makeBibliography()`)
	if err != nil {
		return ret, err
	}
	resExport := result.Export()
	bibEntriesHTML := resExport.([]any)[1].([]any)
	for _, bibEntryHTML := range bibEntriesHTML {
		type inner struct {
			Content string `xml:",innerxml"`
		}
		type outer struct {
			Inners   []inner `xml:"div"`
			InnerXML string  `xml:",innerxml"`
		}
		var o outer
		err := xml.Unmarshal([]byte(bibEntryHTML.(string)), &o)
		if err != nil {
			return ret, err
		}

		if o.Inners == nil {
			ret = append(ret, o.InnerXML)
			continue
		}

		buf := ""
		for i, inner := range o.Inners {
			buf += inner.Content
			if i < len(o.Inners)-1 {
				buf += " "
			}
		}
		ret = append(ret, buf)
	}

	return ret, nil
}
