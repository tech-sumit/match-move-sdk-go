package core_payments

type RemittanceDetailsV1Details struct {
	AmountDeducted float64 `json:"amount_deducted,omitempty"`
	CalculationMode string `json:"calculation_mode,omitempty"`
	AmountSent float64 `json:"amount_sent,omitempty"`
	SendCurrency string `json:"send_currency,omitempty"`
	Fees float64 `json:"fees,omitempty"`
	AmountToReceive float64 `json:"amount_to_receive,omitempty"`
	ReceiveCurrency string `json:"receive_currency,omitempty"`
	ReceiveCountry string `json:"receive_country,omitempty"`
	Provider string `json:"provider,omitempty"`
	DeliveryMethod string `json:"delivery_method,omitempty"`
	PayoutAgent string `json:"payout_agent,omitempty"`
	RoutingParam string `json:"routing_param,omitempty"`
	CustomerFxPercentage string `json:"customer_fx_percentage,omitempty"`
	ExchangeRate float64 `json:"exchange_rate,omitempty"`
	TransactionId string `json:"transaction_id,omitempty"`
	TransactionCode string `json:"transaction_code,omitempty"`
	CrossCurrencyRates float64 `json:"cross_currency_rates,omitempty"`
	AmountBeforeFxPercentage float64 `json:"amount_before_fx_percentage,omitempty"`
	AmountBeforeFxPercentageCcy string `json:"amount_before_fx_percentage_ccy,omitempty"`
	AmountAfterFxPercentage float64 `json:"amount_after_fx_percentage,omitempty"`
	AmountAfterFxPercentageCcy string `json:"amount_after_fx_percentage_ccy,omitempty"`
	SendAmountProvidedByRails float64 `json:"send_amount_provided_by_rails,omitempty"`
	SendAmountProvidedByRailsCcy string `json:"send_amount_provided_by_rails_ccy,omitempty"`
	ActualProviderAmountToSend float64 `json:"actual_provider_amount_to_send,omitempty"`
	ActualProviderAmountToSendCcy string `json:"actual_provider_amount_to_send_ccy,omitempty"`
	ProviderAmountToSend float64 `json:"provider_amount_to_send,omitempty"`
	ProviderAmountToSendCcy string `json:"provider_amount_to_send_ccy,omitempty"`
	ProviderAmountFee float64 `json:"provider_amount_fee,omitempty"`
	ProviderAmountFeeCcy string `json:"provider_amount_fee_ccy,omitempty"`
	BankName string `json:"bank_name,omitempty"`
	BankCode string `json:"bank_code,omitempty"`
	BankBranchName string `json:"bank_branch_name,omitempty"`
	BankBranchCode string `json:"bank_branch_code,omitempty"`
	AmountDetails *RemittanceDetailsV1DetailsAmountDetails `json:"amount_details,omitempty"`
}
