package core_payments

type WalletsStandingInstructionsBody struct {
	BankAccountId string `json:"bank_account_id"`
	Status string `json:"status"`
}
