package core_payments

type RemittanceDetailsV1DetailsAmountDetailsSendDetails struct {
	TotalAmount float64 `json:"total_amount,omitempty"`
	Amount float64 `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
	TotalFee float64 `json:"total_fee,omitempty"`
	RailsFee float64 `json:"rails_fee,omitempty"`
	FixedFee float64 `json:"fixed_fee,omitempty"`
	Tax float64 `json:"tax,omitempty"`
	SubscriptionPricing bool `json:"subscription_pricing,omitempty"`
}
