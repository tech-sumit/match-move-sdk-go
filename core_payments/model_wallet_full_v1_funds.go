package core_payments

type WalletFullV1Funds struct {
	Available *BalanceRootV1 `json:"available"`
	Withholding *BalanceRootV1 `json:"withholding"`
}
