package payment_relay

// Fee breakdown
type TransactionItemFee struct {
	// Basic fee in flat
	FeeFlat float64 `json:"fee_flat,omitempty"`
	// Basic fee in percentage
	FeePercentage float64 `json:"fee_percentage,omitempty"`
	// Tax fee in flat
	TaxFlat float64 `json:"tax_flat,omitempty"`
	// Tax fee in percentage
	TaxPercentage float64 `json:"tax_percentage,omitempty"`
	// Charge fee in flat
	ChargeFlat float64 `json:"charge_flat,omitempty"`
	// Charge fee in percentage
	ChargePercentage float64 `json:"charge_percentage,omitempty"`
	// Service fee in flat
	ServiceFlat float64 `json:"service_flat,omitempty"`
	// Service fee in percentage
	ServicePercentage float64 `json:"service_percentage,omitempty"`
}
