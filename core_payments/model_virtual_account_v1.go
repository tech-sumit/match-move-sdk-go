package core_payments

type VirtualAccountV1 struct {
	AccountNumber float64 `json:"account_number,omitempty"`
	ReversalDestination *BankAccountV1ReversalDestination `json:"reversal_destination,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Status string `json:"status,omitempty"`
}
