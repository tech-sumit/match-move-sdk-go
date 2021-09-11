package core_payments

// Link for attached documents if provided
type RemittanceDetailsV1Attachments struct {
	Rel string `json:"rel,omitempty"`
	Href string `json:"href,omitempty"`
	Method string `json:"method,omitempty"`
}
