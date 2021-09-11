package core_payments

type InlineResponse2004 struct {
	// Status of the card activation process.
	Status string `json:"status,omitempty"`
	Var2fa *InlineResponse20042fa `json:"2fa,omitempty"`
}
