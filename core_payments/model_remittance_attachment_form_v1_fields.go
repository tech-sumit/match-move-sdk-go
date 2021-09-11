package core_payments

type RemittanceAttachmentFormV1Fields struct {
	FieldGroupName string `json:"field_group_name,omitempty"`
	FieldName string `json:"field_name,omitempty"`
	FieldKey string `json:"field_key,omitempty"`
	Type_ string `json:"type,omitempty"`
	Validation *RemittanceAttachmentFormV1Validation `json:"validation,omitempty"`
}
