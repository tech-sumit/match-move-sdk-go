package core_payments

type AuthenticationDocumentV1 struct {
	DocumentId string `json:"document_id,omitempty"`
	DocumentType string `json:"document_type,omitempty"`
	KycType string `json:"kyc_type,omitempty"`
	Link string `json:"link,omitempty"`
}
