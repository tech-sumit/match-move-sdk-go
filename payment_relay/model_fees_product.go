package payment_relay

// Product Details
type FeesProduct struct {
	// Product Code
	Code string `json:"code,omitempty"`
	// Funding Model
	FundingType string `json:"funding_type,omitempty"`
}
