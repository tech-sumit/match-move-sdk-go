package payment_relay

type Response7FeeDetails struct {
	// Flat fee for topup
	FeeFlat string `json:"fee_flat,omitempty"`
	// Percentage fee for topup
	FeePercentage string `json:"fee_percentage,omitempty"`
	// Flat tax
	TaxFlat string `json:"tax_flat,omitempty"`
	// Tax percentage
	TaxPercentage string `json:"tax_percentage,omitempty"`
}
