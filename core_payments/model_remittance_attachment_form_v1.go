package core_payments

type RemittanceAttachmentFormV1 struct {
	Fields []RemittanceAttachmentFormV1Fields `json:"fields,omitempty"`
	RequiredDocuments []string `json:"required_documents,omitempty"`
}
