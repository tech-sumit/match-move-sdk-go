package payment_relay

type Response struct {
	// Provider-specific Merchant Id
	Mid string `json:"mid"`
	// Unique Payment Reference Id
	Ref string `json:"ref"`
	// Payment Acknowledgement Status
	Ack string `json:"ack"`
	// Topup Response Message
	Eod string `json:"eod"`
	Details []ResponseDetails `json:"details"`
}
