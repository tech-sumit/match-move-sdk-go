package core_payments

type VirtualAccountTransaction struct {
	Id int32 `json:"id,omitempty"`
	Amount string `json:"amount,omitempty"`
	Fee string `json:"fee,omitempty"`
	TotalAmount string `json:"total_amount,omitempty"`
	Currency string `json:"currency,omitempty"`
	Status string `json:"status,omitempty"`
	ProcessType string `json:"process_type,omitempty"`
	Details *VirtualAccountTransactionDetails `json:"details,omitempty"`
	UserDetails *OneOfVirtualAccountTransactionUserDetails `json:"user_details,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
