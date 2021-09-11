package core_payments

type InlineResponse20015 struct {
	// Array of documents returned as signed url along with the document type 
	Documents []AnyOfinlineResponse20015DocumentsItems `json:"documents"`
	// Count of documents submitted
	Count int32 `json:"count"`
	// status of the Identity authentication process
	Status string `json:"status"`
	SubStatus string `json:"sub_status,omitempty"`
	// Details passed and used for performing the identity authentications 
	KycDetails *interface{} `json:"kyc_details,omitempty"`
}
