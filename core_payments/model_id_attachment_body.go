package core_payments

type IdAttachmentBody struct {
	// Base64 encoded of the document to attach
	Data string `json:"data,omitempty"`
	// Document type to attach, should be supported on (#get-oauth-consumers-funds-transfers-overseas-attachment-form)
	DocumentType string `json:"document_type,omitempty"`
	// Document name to attach
	DocumentName string `json:"document_name,omitempty"`
}
