package core_payments

type BankAccountV1ReversalDestination struct {
	Id string `json:"id,omitempty"`
	BankAccountNumber string `json:"bank_account_number,omitempty"`
	BankHolderName string `json:"bank_holder_name,omitempty"`
	BankCode string `json:"bank_code,omitempty"`
}
