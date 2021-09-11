package payment_relay

// Fees breakdown
type Response3DetailsFees struct {
	// Currency code [ISO 3166](https://www.iso.org/iso-3166-country-codes.html) alpha-3
	Currency string `json:"currency,omitempty"`
	// Amount entered. Net amount and credit amount will differ depending on the current cppc fee settings.
	Amount float64 `json:"amount,omitempty"`
	// Total fees (provider, tax, charge, service)
	TotalFees float64 `json:"total_fees,omitempty"`
	// Discount offered
	Discount float64 `json:"discount,omitempty"`
	// Total amount to pay
	NetAmount float64 `json:"net_amount,omitempty"`
	// Actual amount to be credited to MM Wallet
	CreditAmount float64 `json:"credit_amount,omitempty"`
	// Mode on how the calculation of credit amount will be. `credit` means credit amount will be the exact amount entered; `net` means the amount entered less fees
	CalculationMode string `json:"calculation_mode,omitempty"`
	// Flag on how the fee will be deducted, whether by `flat` or `percentage`
	Computation string `json:"computation,omitempty"`
	// Fee in Flat
	FeeFlat float64 `json:"fee_flat,omitempty"`
	// Fee in Percentage of the credit amount
	FeePercentage float64 `json:"fee_percentage,omitempty"`
}
