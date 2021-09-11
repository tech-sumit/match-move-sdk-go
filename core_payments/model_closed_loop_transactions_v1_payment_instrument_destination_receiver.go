package core_payments

// This will only show for transaction type = Transferr, show receiver details
type ClosedLoopTransactionsV1PaymentInstrumentDestinationReceiver struct {
	// Receiver Type; email or mobile
	ReceiverType string `json:"receiver_type,omitempty"`
	// Will show if receiver_type = email
	Email string `json:"email,omitempty"`
	// Will show if receiver_type = mobile
	Mobile string `json:"mobile,omitempty"`
}
