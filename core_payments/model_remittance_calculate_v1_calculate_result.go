package core_payments

type RemittanceCalculateV1CalculateResult struct {
	ResultCode string `json:"result_code,omitempty"`
	InitialAmount string `json:"initial_amount,omitempty"`
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
	Quotation *RemittanceQuotationV1 `json:"quotation,omitempty"`
	SubscriptionPricing string `json:"subscription_pricing,omitempty"`
	// For new calculate endpoint, this will not be present on response. Please check on beneficiary_form or sender_form if a particular field is required or not
	RequiredFields string `json:"required_fields,omitempty"`
	AttachmentForm []RemittanceBasicFormV1 `json:"attachment_form,omitempty"`
	SenderForm []RemittanceBasicFormV1 `json:"sender_form,omitempty"`
	BeneficiaryForm []RemittanceBasicFormV1 `json:"beneficiary_form,omitempty"`
}
