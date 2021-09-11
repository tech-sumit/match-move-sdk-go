package core_payments

type RemittanceCalculateV1CalculateResultAmountDetails struct {
	SendDetails *AmountDetailsV2 `json:"send_details,omitempty"`
	ReceiveAmounts *AmountDetailsV2 `json:"receive_amounts,omitempty"`
}
