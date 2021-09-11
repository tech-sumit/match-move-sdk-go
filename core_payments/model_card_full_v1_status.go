package core_payments

type CardFullV1Status struct {
	IsActive bool `json:"is_active,omitempty"`
	Text string `json:"text,omitempty"`
}
