package core_payments

type TransactionPpiPurchaseV1 struct {
	Amount float64 `json:"amount,omitempty"`
	CardTypeCode string `json:"card_type_code,omitempty"`
	Description string `json:"description,omitempty"`
	FundCategoryId string `json:"fund_category_id,omitempty"`
	IsCustomRefId string `json:"is_custom_ref_id,omitempty"`
	Name string `json:"name,omitempty"`
	ChargeFee string `json:"charge_fee,omitempty"`
	TransactionCurrency string `json:"transaction_currency,omitempty"`
	Reason string `json:"reason,omitempty"`
	Tt string `json:"tt,omitempty"`
	MerchantName string `json:"merchant_name,omitempty"`
	MerchantCategoryCode string `json:"merchant_category_code,omitempty"`
	MerchantLocation string `json:"merchant_location,omitempty"`
	TerminalId string `json:"terminal_id,omitempty"`
	OriginalTransactionCurrency string `json:"original_transaction_currency,omitempty"`
	OriginalTransactionAmount string `json:"original_transaction_amount,omitempty"`
	TotalFeesCharged string `json:"total_fees_charged,omitempty"`
	TotalFeesCurrency string `json:"total_fees_currency,omitempty"`
	TotalTransactionAmount string `json:"total_transaction_amount,omitempty"`
	TotalTransactionCurrency string `json:"total_transaction_currency,omitempty"`
	TransactionMode string `json:"transaction_mode,omitempty"`
	TransactionNetwork string `json:"transaction_network,omitempty"`
	AdditionalDetails *TransactionFundsDebitreversalV1AdditionalDetails `json:"additional_details,omitempty"`
	DateReversed string `json:"date_reversed,omitempty"`
	PaymentRef string `json:"payment_ref,omitempty"`
	Pp string `json:"pp,omitempty"`
	ProviderRefId string `json:"provider_ref_id,omitempty"`
	Vat string `json:"vat,omitempty"`
	CardFee float64 `json:"card_fee,omitempty"`
}
