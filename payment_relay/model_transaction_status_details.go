package payment_relay

// Topup Details
type TransactionStatusDetails struct {
	Fees *TransactionStatusDetailsFees `json:"fees,omitempty"`
	User *TransactionStatusDetailsUser `json:"user,omitempty"`
	Source *TransactionStatusDetailsSource `json:"source,omitempty"`
	// Payment Provider Status Inquiry
	Provider *interface{} `json:"provider,omitempty"`
}
