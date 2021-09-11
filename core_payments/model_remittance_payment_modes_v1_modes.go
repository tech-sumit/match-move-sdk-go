package core_payments

type RemittancePaymentModesV1Modes struct {
	// Payment mode name
	Name string `json:"name,omitempty"`
	// Payment mode
	Mode string `json:"mode,omitempty"`
	Providers []RemittancePaymentModesV1Providers `json:"providers,omitempty"`
}
