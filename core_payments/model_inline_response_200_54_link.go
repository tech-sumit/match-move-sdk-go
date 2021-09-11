package core_payments

type InlineResponse20054Link struct {
	// Identifier if next or previous page
	Rel string `json:"rel,omitempty"`
	// Pagination link
	Href string `json:"href,omitempty"`
	// Pagination method
	Method string `json:"method,omitempty"`
}
