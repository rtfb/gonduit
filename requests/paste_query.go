package requests

// PasteQueryRequest represents a request to paste.query.
type PasteQueryRequest struct {
	IDs         []uint64 `json:"ids"`         // optional
	PHIDs       []string `json:"phids"`       // optional
	AuthorPHIDs []string `json:"authorPHIDs"` // optional
	Offset      uint64   `json:"after"`       // optional
	Limit       uint64   `json:"limit"`       // optional
	Request
}
