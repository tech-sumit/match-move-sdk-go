package core_payments

type RemittanceDetailsV1 struct {
	// Remittance Hash ID
	RefId string `json:"ref_id,omitempty"`
	Type_ string `json:"type,omitempty"`
	ClientKey string `json:"client_key,omitempty"`
	StatusCode string `json:"status_code,omitempty"`
	StatusText string `json:"status_text,omitempty"`
	Remarks string `json:"remarks,omitempty"`
	Date string `json:"date,omitempty"`
	DateExpiry string `json:"date_expiry,omitempty"`
	// Document Reference; will only show if has value
	DocumentReferenceNumber string `json:"document_reference_number,omitempty"`
	// Client generated reference; will only show if has value
	ClientRefId string `json:"client_ref_id,omitempty"`
	TransactionCode string `json:"transaction_code,omitempty"`
	Sender *RemittanceDetailsV1Sender `json:"sender,omitempty"`
	Receiver *AnyOfRemittanceDetailsV1Receiver `json:"receiver,omitempty"`
	Details *RemittanceDetailsV1Details `json:"details,omitempty"`
	Confirm string `json:"confirm,omitempty"`
	Attachments *RemittanceDetailsV1Attachments `json:"attachments,omitempty"`
}
