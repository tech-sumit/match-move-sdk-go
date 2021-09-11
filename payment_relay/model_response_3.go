package payment_relay

type Response3 struct {
	// Unique Payment Token Reference
	Ref string `json:"ref"`
	// Payment Acknowledgement Status
	Ack string `json:"ack"`
	// Topup Response Message
	Eod string `json:"eod"`
	Details *Response3Details `json:"details"`
}
