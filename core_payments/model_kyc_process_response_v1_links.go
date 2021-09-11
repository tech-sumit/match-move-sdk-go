package core_payments

type KycProcessResponseV1Links struct {
	// HTTP Link method
	Method string `json:"method,omitempty"`
	// HTTP Link url (for the SDK)
	Href string `json:"href,omitempty"`
	// HTTP Link relationship
	Rel string `json:"rel,omitempty"`
}
