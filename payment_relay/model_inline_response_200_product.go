package payment_relay

type InlineResponse200Product struct {
	// Product Code
	Code string `json:"code,omitempty"`
	// Funding Model
	FundingType string `json:"funding_type,omitempty"`
}
