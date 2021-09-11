package core_payments

type TransactionCategoriesRulesRequestV1 struct {
	// Event code from event_codes table
	EventCodes string `json:"event_codes"`
	// Transcation channel like atm,pos,ecom
	Channel string `json:"channel,omitempty"`
	// 3 Digits iso defined country code number like 356
	AcquiringCountryCode string `json:"acquiring_country_code,omitempty"`
	// currency code like USD
	Currency string `json:"currency,omitempty"`
	// Transcation network like visa,mastercard
	Network string `json:"network,omitempty"`
	// terminal_id defined in transcation rules table
	TerminalId string `json:"terminal_id,omitempty"`
	// merchant_id defined in transcation rules table
	MerchantId string `json:"merchant_id,omitempty"`
	// Point of service entry mode like MANUALLY_KEYED
	PosEntryModeValue string `json:"pos_entry_mode_value,omitempty"`
	// pp value that goes into pp cosnumer table
	Pp string `json:"pp,omitempty"`
	// merchant_category_code defined in transcation rules table
	MerchantCategoryCode string `json:"merchant_category_code,omitempty"`
}
