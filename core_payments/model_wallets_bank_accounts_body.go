package core_payments

type WalletsBankAccountsBody struct {
	BankAccountNumber string `json:"bank_account_number"`
	BankHolderName string `json:"bank_holder_name"`
	BankCode string `json:"bank_code"`
}
