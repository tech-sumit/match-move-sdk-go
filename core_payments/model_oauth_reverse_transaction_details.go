package core_payments

type OauthReverseTransactionDetails struct {
	Amount float64 `json:"amount,omitempty"`
	Fee float64 `json:"fee,omitempty"`
	Total int32 `json:"total,omitempty"`
	Tt string `json:"tt,omitempty"`
	Pp string `json:"pp,omitempty"`
	RefId string `json:"ref_id,omitempty"`
	PaymentRef string `json:"payment_ref,omitempty"`
	IsCustomRefId string `json:"is_custom_ref_id,omitempty"`
}
