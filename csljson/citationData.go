package csljson

// CitationData contains the citation data for a CSL JSON citation.
type CitationData struct {
	// Items is an array of items, each representing a single source of information.
	Items []Item `json:"items,omitempty"`
	// Properties is an object that contains properties related to the citation, such as the note index.
	Properties struct {
		// NoteIndex is the index of the citation in the list of notes.
		NoteIndex int `json:"noteIndex,omitempty"`
	}
}
