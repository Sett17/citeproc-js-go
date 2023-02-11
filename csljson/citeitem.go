package csljson

// CiteItem represents a single source of information in a CSL JSON citation.
type CiteItem struct {
	// ID is a unique identifier for the item.
	ID string `json:"id"`
	// Locator is a string that indicates the specific location within the item, such as a page number.
	Locator string `json:"locator"`
	// SuppressAuthor is a flag that indicates whether the author's name should be suppressed.
	SuppressAuthor bool `json:"suppress-author"`
	// AuthorOnly is a flag that indicates whether only the author's name should be included in the citation.
	AuthorOnly bool `json:"author-only"`
	// Prefix is a string that is added before the citation.
	Prefix string `json:"prefix"`
	// Suffix is a string that is added after the citation.
	Suffix string `json:"suffix"`

	// Only with makeCitationCluster():
	// Position is the position of the citation item in the citation.
	Position int `json:"position"`
	// NearNote is a flag that indicates whether the citation is near a note.
	NearNote bool `json:"near-note"`
}

// CiteItemFromItem creates a CiteItem from an Item.
func CiteItemFromItem(i Item) CiteItem {
	return CiteItem{
		ID: i.ID,
	}
}
