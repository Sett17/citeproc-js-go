package csljson

// Citation represents a CSL JSON citation.
type Citation struct {
	// CitationItems is an array of citation items, each representing a single source of information.
	CitationItems []CiteItem `json:"citationItems"`
	// Properties is an object that contains properties related to the citation, such as the note index.
	Properties struct {
		// NoteIndex is the index of the citation in the list of notes.
		NoteIndex int `json:"noteIndex"`
	} `json:"properties"`

	// CitationID is an identifier for the citation.
	CitationID any `json:"citationID"`
	// SortedItems is an array of citation items that have been sorted according to the citation style.
	SortedItems []any `json:"sortedItems"`
}
