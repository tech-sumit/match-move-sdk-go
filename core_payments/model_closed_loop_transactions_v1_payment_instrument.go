package core_payments

type ClosedLoopTransactionsV1PaymentInstrument struct {
	Source *ClosedLoopTransactionsV1PaymentInstrumentSource `json:"source,omitempty"`
	Destination *ClosedLoopTransactionsV1PaymentInstrumentDestination `json:"destination,omitempty"`
}
