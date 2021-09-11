package payment_relay

type Response2 struct {
	// Is request `approved` or `rejected`
	Status string `json:"status"`
}
