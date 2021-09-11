package payment_relay

type Response4Product struct {
	// Product Code
	Code string `json:"code,omitempty"`
	// Funding Model
	FundingType string `json:"funding_type,omitempty"`
}
