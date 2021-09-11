package core_payments

type AmountDetailsV1 struct {
	Amount string `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
	Country *AmountDetailsV1Country `json:"country,omitempty"`
}
