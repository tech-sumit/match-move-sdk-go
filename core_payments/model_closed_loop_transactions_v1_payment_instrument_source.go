package core_payments

type ClosedLoopTransactionsV1PaymentInstrumentSource struct {
	// Source Type; wallet or card
	SourceType string `json:"source_type,omitempty"`
	// Wallet or Card Hash ID
	HashId string `json:"hash_id,omitempty"`
	// Wallet or Card Proxy Number
	ProxyNumber string `json:"proxy_number,omitempty"`
	// Wallet or Card Masked Card Number
	Number string `json:"number,omitempty"`
	// Wallet or Card Status
	State string `json:"state,omitempty"`
	Sender *ClosedLoopTransactionsV1PaymentInstrumentSourceSender `json:"sender,omitempty"`
}
