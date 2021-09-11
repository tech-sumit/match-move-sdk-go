package core_payments

type RemittanceAgentV1Fees struct {
	// Fixed fee value to deduct
	FixedFee string `json:"fixed_fee"`
	// FX Markup Percentage
	FxMarkup string `json:"fx_markup"`
}
