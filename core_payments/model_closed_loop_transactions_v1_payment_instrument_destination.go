package core_payments

type ClosedLoopTransactionsV1PaymentInstrumentDestination struct {
	// Source Type; wallet or card
	DestinationType string `json:"destination_type,omitempty"`
	// Wallet or Card Hash ID
	HashId string `json:"hash_id,omitempty"`
	// Wallet or Card Proxy Number
	ProxyNumber string `json:"proxy_number,omitempty"`
	// Wallet or Card Masked Card Number
	Number string `json:"number,omitempty"`
	// Wallet or Card Status
	State string `json:"state,omitempty"`
	Receiver *ClosedLoopTransactionsV1PaymentInstrumentDestinationReceiver `json:"receiver,omitempty"`
}
