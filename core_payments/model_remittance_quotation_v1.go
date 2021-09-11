package core_payments

type RemittanceQuotationV1 struct {
	// Quotation ID from Remit Service
	Id string `json:"id"`
	// External ID from Rails
	ExternalId string `json:"external_id"`
	// Quotation ID from Rails
	QuotationId string `json:"quotation_id"`
	// Payout Agent Identifier
	PayoutAgent string `json:"payout_agent"`
	// Calculation Mode
	CalculationMode string `json:"calculation_mode"`
	// Transaction Type - C2C or B2B
	TransactionType string `json:"transaction_type"`
	Initial *AmountDetailsV1 `json:"initial"`
	Send *AmountDetailsV1 `json:"send"`
	Receive *AmountDetailsV1 `json:"receive"`
	Fee *AmountDetailsV1 `json:"fee"`
	// Calculated Exchange Rate after fx markup
	ExchangeRate string `json:"exchange_rate"`
	// Exchange Rate from Rails
	ProviderExchangeRate string `json:"provider_exchange_rate"`
	// Quotation Status
	Status string `json:"status"`
	// Quotation Expiry (after an hour it was created for T2)
	DateExpiry string `json:"date_expiry"`
	// Quotation Date Created
	DateCreated string `json:"date_created"`
	// Set to true if subscription pricing was applied, else false
	SubscriptionPricing bool `json:"subscription_pricing"`
}
