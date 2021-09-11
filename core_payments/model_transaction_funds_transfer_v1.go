package core_payments

type TransactionFundsTransferV1 struct {
	// Unique ID for the transaction
	Id string `json:"id,omitempty"`
	Sender *TransactionFundsTransferV1Sender `json:"sender,omitempty"`
	Recipient *TransactionFundsTransferV1Sender `json:"recipient,omitempty"`
	// Message passed for the transaction
	Message string `json:"message,omitempty"`
	Status string `json:"status,omitempty"`
	Amount float64 `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
	Date *TransactionFundsTransferV1Date `json:"date,omitempty"`
}
