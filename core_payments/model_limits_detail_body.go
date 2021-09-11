package core_payments

type LimitsDetailBody struct {
	CardHashId string `json:"card_hash_id,omitempty"`
	UserHashId string `json:"user_hash_id,omitempty"`
	TransactionAmount float64 `json:"transaction_amount,omitempty"`
	Amount float64 `json:"amount,omitempty"`
	CardCurrency string `json:"card_currency,omitempty"`
	TransactionCurrency string `json:"transaction_currency,omitempty"`
	Channel string `json:"channel,omitempty"`
	PosEntryModeValue string `json:"pos_entry_mode_value,omitempty"`
	AcquiringCountryCode string `json:"acquiring_country_code,omitempty"`
	Network string `json:"network,omitempty"`
	MerchantCategoryCode string `json:"merchant_category_code,omitempty"`
	MerchantId string `json:"merchant_id,omitempty"`
	TerminalId string `json:"terminal_id,omitempty"`
}
