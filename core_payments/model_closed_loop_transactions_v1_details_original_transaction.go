package core_payments

// This will only show for reversal and refund transaction type
type ClosedLoopTransactionsV1DetailsOriginalTransaction struct {
	// Transaction Type
	TransactionType string `json:"transaction_type,omitempty"`
	// Transaction Reference Hash
	RefHash string `json:"ref_hash,omitempty"`
	// Transaction Currency
	Currency string `json:"currency,omitempty"`
	// Transaction Amount
	Amount float32 `json:"amount,omitempty"`
}
