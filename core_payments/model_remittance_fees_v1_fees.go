package core_payments

type RemittanceFeesV1Fees struct {
	Id string `json:"id,omitempty"`
	CalculationMode string `json:"calculation_mode,omitempty"`
	PaymentMode string `json:"payment_mode,omitempty"`
	PaymentModeDescription string `json:"payment_mode_description,omitempty"`
	RoutingParam string `json:"routing_param,omitempty"`
	RoutingType string `json:"routing_type,omitempty"`
	ProviderCurrency string `json:"provider_currency,omitempty"`
	SendCurrency string `json:"send_currency,omitempty"`
	ReceiveCountry string `json:"receive_country,omitempty"`
	ReceiveCurrency string `json:"receive_currency,omitempty"`
	PartnerName string `json:"partner_name,omitempty"`
	Precision string `json:"precision,omitempty"`
	ExchangeRate float64 `json:"exchange_rate,omitempty"`
	FixedFee string `json:"fixed_fee,omitempty"`
	PayerDetails *RemittanceFeesV1PayerDetails `json:"payer_details,omitempty"`
	FullDetails *RemittanceFeesV1FullDetails `json:"full_details,omitempty"`
	ProviderExchangeRate string `json:"provider_exchange_rate,omitempty"`
	CrossRateDetails *RemittanceFeesV1CrossRateDetails `json:"cross_rate_details,omitempty"`
	FxMarkupPercentage string `json:"fx_markup_percentage,omitempty"`
	FxRate string `json:"fx_rate,omitempty"`
	PricingsId string `json:"pricings_id,omitempty"`
	CurrencyFxRate float64 `json:"currency_fx_rate,omitempty"`
	ReverseCurrencyFxRate float64 `json:"reverse_currency_fx_rate,omitempty"`
}
