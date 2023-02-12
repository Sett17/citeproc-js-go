package csljson

//// Item is a struct that represents an item in a citation information.
//type Item struct {
//	// ID is the identifier of the item.
//	ID string `json:"id"`
//	// Type is the type of the item.
//	Type string `json:"type"`
//	// Abstract is a brief summary of the item.
//	Abstract string `json:"abstract"`
//	// Accessed is a struct that contains the date when the item was accessed.
//	Accessed struct {
//		// DateParts is a two-dimensional array that contains the date parts in integer form.
//		DateParts [][]any `json:"date-parts"` // is numbers and string in regular csl-json, better csl-json only numbers
//	} `json:"accessed"`
//	// Archive is the name of the archive where the item is stored.
//	Archive string `json:"archive"`
//	// ArchiveLocation is the location of the item in the archive.
//	ArchiveLocation string `json:"archive_location"`
//	// Author is an array of structs that contain the author's family name and given name.
//	Author []struct {
//		// Family is the author's family name.
//		Family string `json:"family"`
//		// Given is the author's given name.
//		Given string `json:"given"`
//	} `json:"author"`
//	// CallNumber is the call number of the item in a library.
//	CallNumber string `json:"call-number"`
//	// CitationKey is a unique identifier for the item used in a citation.
//	CitationKey string `json:"citation-key"`
//	// Dimensions is the physical dimensions of the item.
//	Dimensions string `json:"dimensions"`
//	// Issued is a struct that contains the date when the item was issued.
//	Issued struct {
//		// Literal is a string representation of the date when the item was issued.
//		Literal string `json:"literal"`
//	} `json:"issued"`
//	// Language is the language of the item.
//	Language string `json:"language"`
//	// License is the license under which the item is available.
//	License string `json:"license"`
//	// Medium is the medium in which the item is stored.
//	Medium string `json:"medium"`
//	// Note is any additional notes about the item.
//	Note string `json:"note"`
//	// Source is the source of the item.
//	Source string `json:"source"`
//	// Title is the title of the item.
//	Title string `json:"title"`
//	// TitleShort is a shortened version of the title of the item.
//	TitleShort string `json:"title-short"`
//	// URL is the URL where the item can be accessed.
//	URL string `json:"URL"`
//	// CollectionTitle is the title of the collection that the item is part of.
//	CollectionTitle string `json:"collection-title"`
//	// EventPlace is the place where an event associated with the item took place.
//	EventPlace string `json:"event-place"`
//	// Isbn is the ISBN of the item.
//	Isbn string `json:"ISBN"`
//	// NumberOfVolumes is the number of volumes in the item.
//	NumberOfVolumes string `json:"number-of-volumes"`
//	// Publisher holds information about the publisher of the item
//	Publisher string `json:"publisher"`
//	// PublisherPlace holds information about the location of the publisher
//	PublisherPlace string `json:"publisher-place"`
//	// Volume holds information about the volume of the item
//	Volume string `json:"volume"`
//	// Authority holds information about the authority responsible for the item
//	Authority string `json:"authority"`
//	// ChapterNumber holds information about the chapter number of the item
//	ChapterNumber string `json:"chapter-number"`
//	// ContainerTitle holds information about the title of the container in which the item is published
//	ContainerTitle string `json:"container-title"`
//	// Number holds information about the number of the item
//	Number string `json:"number"`
//	// Page holds information about the page number of the item
//	Page string `json:"page"`
//	// References holds information about the references cited in the item
//	References string `json:"references"`
//	// Section holds information about the section of the item
//	Section string `json:"section"`
//	// Genre holds information about the genre of the item
//	Genre string `json:"genre"`
//	// CollectionNumber holds information about the collection number of the item
//	CollectionNumber string `json:"collection-number"`
//	// Edition holds information about the edition of the item
//	Edition string `json:"edition"`
//	// NumberOfPages holds information about the number of pages of the item
//	NumberOfPages string `json:"number-of-pages"`
//	// Doi holds information about the DOI of the item
//	Doi string `json:"DOI"`
//	// EventTitle holds information about the title of the event in which the item was published
//	EventTitle string `json:"event-title"`
//	// Director holds information about the director(s) of the item
//	Director []struct {
//		Family string `json:"family"`
//		Given  string `json:"given"`
//	} `json:"director"`
//	// Contributor holds information about the contributor(s) of the item
//	Contributor []struct {
//		Family string `json:"family"`
//		Given  string `json:"given"`
//	} `json:"contributor"`
//	// ContainerTitleShort holds information about the short title of the container in which the item is published
//	ContainerTitleShort string `json:"container-title-short"`
//	// Issn holds information about the ISSN of the item
//	Issn string `json:"ISSN"`
//	// Issue holds information about the issue of the item
//	Issue string `json:"issue"`
//	// Scale holds information about the scale of the item
//	Scale string `json:"scale"`
//	// Submitted holds information about when the item was submitted
//	Submitted struct {
//		Literal string `json:"literal"`
//	} `json:"submitted"`
//	// Version holds information about the version of the item
//	Version string `json:"version"`
//}

type Item struct {
	ID             string `json:"id"`
	Type           string `json:"type"`
	ContainerTitle string `json:"container-title,omitempty"`
	Language       string `json:"language,omitempty"`
	Title          string `json:"title"`
	URL            string `json:"URL,omitempty"`
	Accessed       struct {
		DateParts [][]interface{} `json:"date-parts"`
	} `json:"accessed,omitempty"`
	Abstract string `json:"abstract,omitempty"`
	License  string `json:"license,omitempty"`
	Note     string `json:"note,omitempty"`
	Issued   struct {
		DateParts [][]interface{} `json:"date-parts"`
	} `json:"issued,omitempty"`
	Genre  string `json:"genre,omitempty"`
	Source string `json:"source,omitempty"`
	Author []struct {
		Family string `json:"family"`
		Given  string `json:"given"`
	} `json:"author,omitempty"`
	TitleShort string `json:"title-short,omitempty"`
	Number     string `json:"number,omitempty"`
	Publisher  string `json:"publisher,omitempty"`
}
