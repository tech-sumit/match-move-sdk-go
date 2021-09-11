package core_payments

type ClosedLoopTransactionsV1Details struct {
	// Debit/Credit indicator
	Type_ string `json:"type,omitempty"`
	// Transaction Reference Identifier
	RefHash string `json:"ref_hash,omitempty"`
	// Transaction Currency
	Currency string `json:"currency,omitempty"`
	// Transaction Amount
	Amount float32 `json:"amount,omitempty"`
	// Wallet or card current balance based on the transaction
	Balance float32 `json:"balance,omitempty"`
	// Transaction Status
	Status string `json:"status,omitempty"`
	// Transaction Date Added
	Date string `json:"date,omitempty"`
	// Additional information of a particular transaction, values can be different depending on transaction type
	AdditionalInformation string `json:"additional_information,omitempty"`
	OriginalTransaction *ClosedLoopTransactionsV1DetailsOriginalTransaction `json:"original_transaction,omitempty"`
}
