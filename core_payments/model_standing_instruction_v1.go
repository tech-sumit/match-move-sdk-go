package core_payments

type StandingInstructionV1 struct {
	Id string `json:"id,omitempty"`
	BankAccountId string `json:"bank_account_id,omitempty"`
	Number float64 `json:"number,omitempty"`
	MaximumDebitableAmount string `json:"maximum_debitable_amount,omitempty"`
	DateApplication string `json:"date_application,omitempty"`
	DateExpiry string `json:"date_expiry,omitempty"`
	ConsentDetails *ConsentV1 `json:"consent_details,omitempty"`
	Status string `json:"status,omitempty"`
}
