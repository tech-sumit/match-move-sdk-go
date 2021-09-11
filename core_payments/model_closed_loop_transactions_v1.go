package core_payments

type ClosedLoopTransactionsV1 struct {
	// User Hash ID
	Id string `json:"id,omitempty"`
	// User Email
	Email string `json:"email,omitempty"`
	// User Full Mobile
	Mobile string `json:"mobile,omitempty"`
	// Transaction Type
	TransactionType string `json:"transaction_type,omitempty"`
	PaymentInstrument *ClosedLoopTransactionsV1PaymentInstrument `json:"payment_instrument,omitempty"`
	Details *ClosedLoopTransactionsV1Details `json:"details,omitempty"`
}
