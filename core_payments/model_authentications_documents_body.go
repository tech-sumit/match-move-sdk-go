package core_payments

type AuthenticationsDocumentsBody struct {
	// base64 of a file content. Accepted doc types [JPG,PNG,PDF]. Max size of file supported is 8MB.
	Data string `json:"data,omitempty"`
	DocumentType string `json:"document_type,omitempty"`
	KycType string `json:"kyc_type,omitempty"`
}
