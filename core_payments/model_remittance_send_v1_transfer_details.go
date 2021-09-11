package core_payments

type RemittanceSendV1TransferDetails struct {
	CalculationMode string `json:"calculation_mode,omitempty"`
	InitialAmount string `json:"initial_amount,omitempty"`
	ReferenceNumber string `json:"reference_number,omitempty"`
	AmountToSend string `json:"amount_to_send,omitempty"`
	AmountToSendCurrency string `json:"amount_to_send_currency,omitempty"`
	AmountToReceive string `json:"amount_to_receive,omitempty"`
	AmountToReceiveCurrency string `json:"amount_to_receive_currency,omitempty"`
	ReceiveCurrency string `json:"receive_currency,omitempty"`
	ProviderExchangeRate string `json:"provider_exchange_rate,omitempty"`
	ExchangeRate string `json:"exchange_rate,omitempty"`
	FixedFee string `json:"fixed_fee,omitempty"`
	Commission string `json:"commission,omitempty"`
	TotalAmount string `json:"total_amount,omitempty"`
	AmountToBeDeducted string `json:"amount_to_be_deducted,omitempty"`
	AmountDetails *RemittanceCalculateV1CalculateResultAmountDetails `json:"amount_details,omitempty"`
}
