package core_payments

type OauthReverseTransaction struct {
	Id string `json:"id,omitempty"`
	Type_ string `json:"type,omitempty"`
	Amount string `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
	Status string `json:"status,omitempty"`
	Indicator string `json:"indicator,omitempty"`
	CardNumber string `json:"card_number,omitempty"`
	SenderVpa string `json:"sender_vpa,omitempty"`
	ReceiverVpa string `json:"receiver_vpa,omitempty"`
	Details *OauthReverseTransactionDetails `json:"details,omitempty"`
	Date *OauthReverseTransactionDate `json:"date,omitempty"`
}
