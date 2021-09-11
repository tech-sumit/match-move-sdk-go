package payment_relay

type Response1 struct {
	// Unique Payment Reference Id
	Ref string `json:"ref"`
	// Payment Acknowledgement Status
	Ack string `json:"ack"`
	// Topup Response Message
	Eod string `json:"eod"`
}
