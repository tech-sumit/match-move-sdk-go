package core_payments

type VirtualAccountTransactionDetails struct {
	BankRef string `json:"bank_ref,omitempty"`
	CustomerRef string `json:"customer_ref,omitempty"`
	JournalId string `json:"journal_id,omitempty"`
	TopupType string `json:"topup_type,omitempty"`
	VirtualAccountNumber string `json:"virtual_account_number,omitempty"`
}
