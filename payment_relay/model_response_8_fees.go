package payment_relay

// Fees breakdown
type Response8Fees struct {
	// Computation type
	Computation string `json:"computation,omitempty"`
	// Flat Fee
	FeeFlat float64 `json:"fee_flat,omitempty"`
	// Fee percentage
	FeePercentage float64 `json:"fee_percentage,omitempty"`
}
