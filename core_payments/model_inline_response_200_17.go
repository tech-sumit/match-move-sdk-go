package core_payments

type InlineResponse20017 struct {
	// Array of documents returned as signed url along with the document type 
	Documents []AnyOfinlineResponse20017DocumentsItems `json:"documents,omitempty"`
	// Count of documents submitted
	Count int32 `json:"count,omitempty"`
	// status of the Identity authentication process
	Status string `json:"status,omitempty"`
}
