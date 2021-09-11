package payment_relay

// Fees breakdown
type Response3DetailsTopup struct {
	// Currency code [ISO 3166](https://www.iso.org/iso-3166-country-codes.html) alpha-3
	Currency string `json:"currency,omitempty"`
	// Actual amount to be credited to MM Wallet
	Amount float64 `json:"amount,omitempty"`
}
