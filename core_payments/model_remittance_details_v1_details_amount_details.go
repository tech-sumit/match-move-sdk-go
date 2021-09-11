package core_payments

type RemittanceDetailsV1DetailsAmountDetails struct {
	SendDetails *RemittanceDetailsV1DetailsAmountDetailsSendDetails `json:"send_details,omitempty"`
	ReceiveAmounts *RemittanceDetailsV1DetailsAmountDetailsReceiveAmounts `json:"receive_amounts,omitempty"`
}
