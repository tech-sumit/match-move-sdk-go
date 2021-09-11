package core_payments

type FundTransferV1 struct {
	Id string `json:"id,omitempty"`
	ClientRefId string `json:"client_ref_id,omitempty"`
	UniquePaymentId string `json:"unique_payment_id,omitempty"`
	Currency string `json:"currency,omitempty"`
	PurposeOfTransfer string `json:"purpose_of_transfer,omitempty"`
	TransferType string `json:"transfer_type,omitempty"`
	Status string `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
	BankAccount *BankAccountV1ReversalDestination `json:"bank_account,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
