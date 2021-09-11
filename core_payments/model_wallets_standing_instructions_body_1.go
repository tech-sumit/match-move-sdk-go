package core_payments

type WalletsStandingInstructionsBody1 struct {
	BankAccountId string `json:"bank_account_id"`
	Number float64 `json:"number"`
	MaximumDebitableAmount string `json:"maximum_debitable_amount"`
	// Example 2020-01-01 12:00:00
	DateApplication string `json:"date_application"`
	// Example 2020-01-01
	DateExpiry string `json:"date_expiry"`
	Status string `json:"status"`
	ConsentFullName string `json:"consent_full_name"`
	ConsentNumbers string `json:"consent_numbers"`
	ConsentSign string `json:"consent_sign"`
}
