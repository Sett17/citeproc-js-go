package citeproc_js_go

import (
	_ "embed"
	"encoding/json"
	"github.com/sett17/citeproc-js-go/csljson"
	"testing"
)

//go:embed citations.json
var citationsFile []byte

func TestDev(t *testing.T) {
	s := NewSession()
	err := s.Init()
	if err != nil {
		t.Fatal(err)
	}

	//citations := make(map[string]csljson.Item)
	citsList := make([]csljson.Item, 0)
	err = json.Unmarshal(citationsFile, &citsList)
	if err != nil {
		t.Fatal(err)
	}

	err = s.AddCitation(citsList...)
	if err != nil {
		t.Fatal(err)
	}

	res1, err := s.ProcessCitationCluster(citsList[0:3]...)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("res1: %+v", res1)

	res2, err := s.ProcessCitationCluster(citsList[1], citsList[4], citsList[6], citsList[7])
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("res1: %+v", res2)

	bib, err := s.MakeBibliography()
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range bib {
		t.Logf("%s", entry)
	}

}
