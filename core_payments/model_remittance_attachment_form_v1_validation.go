package core_payments

type RemittanceAttachmentFormV1Validation struct {
	IsRequired string `json:"is_required"`
	MimeType string `json:"mime_type,omitempty"`
}
