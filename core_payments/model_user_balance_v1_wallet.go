package core_payments

type UserBalanceV1Wallet struct {
	Available *UserBalanceV1WalletAvailable `json:"available,omitempty"`
	Withholding *UserBalanceV1WalletWithholding `json:"withholding,omitempty"`
}
