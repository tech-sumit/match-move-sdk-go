package core_payments

type UserNameV1 struct {
	// First or given name
	First string `json:"first,omitempty"`
	// Middle name
	Middle string `json:"middle,omitempty"`
	// Last or family name
	Last string `json:"last,omitempty"`
	// Preferred Name or Name on Card
	Preferred string `json:"preferred,omitempty"`
}
