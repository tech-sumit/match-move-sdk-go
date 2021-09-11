package payment_relay

// User Bank Details
type VirtualAddressBank struct {
	// Bank Name
	Name string `json:"name,omitempty"`
	// Bank Account Number
	AccountNumber string `json:"account_number,omitempty"`
	// Bank Account Type
	AccountType string `json:"account_type,omitempty"`
	// Bank IFSC code
	Ifsc string `json:"ifsc,omitempty"`
}
