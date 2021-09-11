package core_payments

type RemittanceAttachmentV1 struct {
	Id string `json:"id,omitempty"`
	DocumentType string `json:"document_type,omitempty"`
	ContentType string `json:"content_type,omitempty"`
	DocumentName string `json:"document_name,omitempty"`
	Link string `json:"link,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}
