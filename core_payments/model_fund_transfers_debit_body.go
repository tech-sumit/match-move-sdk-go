package core_payments

type FundTransfersDebitBody struct {
	BankAccountId string `json:"bank_account_id"`
	ClientRefId string `json:"client_ref_id"`
	Amount float64 `json:"amount"`
	Currency string `json:"currency"`
	PurposeOfTransfer string `json:"purpose_of_transfer"`
}
