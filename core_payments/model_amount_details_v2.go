package core_payments

type AmountDetailsV2 struct {
	TotalAmount string `json:"total_amount,omitempty"`
	Amount string `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
	TotalFee string `json:"total_fee,omitempty"`
	RailsFee string `json:"rails_fee,omitempty"`
	FixedFee string `json:"fixed_fee,omitempty"`
	Tax string `json:"tax,omitempty"`
	SubscriptionPricing string `json:"subscription_pricing,omitempty"`
}
