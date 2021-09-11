package payment_relay

type Response5Links struct {
	// Self-link to href value
	Rel string `json:"rel,omitempty"`
	// Payment Gateway URL
	Href string `json:"href,omitempty"`
	// Dynamic values returned from the payment gateway
	Parameters *interface{} `json:"parameters,omitempty"`
}
