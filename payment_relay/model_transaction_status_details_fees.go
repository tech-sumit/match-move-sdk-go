package payment_relay

// Fees breakdown
type TransactionStatusDetailsFees struct {
	// Currency code [ISO 3166](https://www.iso.org/iso-3166-country-codes.html) alpha-3
	Currency string `json:"currency,omitempty"`
	// Amount entered
	Amount float64 `json:"amount,omitempty"`
	// Total fees (provider, tax, charge, service)
	TotalFees float64 `json:"total_fees,omitempty"`
	// Discount offered
	Discount float64 `json:"discount,omitempty"`
	// Total amount to pay
	NetAmount float64 `json:"net_amount,omitempty"`
	// Actual amount to be credited to MM Wallet
	CreditAmount float64 `json:"credit_amount,omitempty"`
	// flat fee
	FeeFlat float64 `json:"fee_flat,omitempty"`
	// percentage fee
	FeePercentage float64 `json:"fee_percentage,omitempty"`
}
