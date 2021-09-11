package core_payments

type FundTransfersCreditBody struct {
	BankAccountId string `json:"bank_account_id"`
	ClientRefId string `json:"client_ref_id"`
	Amount float64 `json:"amount"`
	// The currency of the application.
	Currency string `json:"currency"`
	PurposeOfTransfer string `json:"purpose_of_transfer"`
}
