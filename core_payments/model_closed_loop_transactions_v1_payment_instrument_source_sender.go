package core_payments

// This will only show for transaction type = Transfer, show sender details
type ClosedLoopTransactionsV1PaymentInstrumentSourceSender struct {
	// Sender Type; email or mobile
	SenderType string `json:"sender_type,omitempty"`
	// Will show if sender_type = email
	Email string `json:"email,omitempty"`
	// Will show if sender_type = mobile
	Mobile string `json:"mobile,omitempty"`
}
