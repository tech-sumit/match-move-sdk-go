package core_payments

type UserBalanceV1WalletWithholding struct {
	Currency string `json:"currency,omitempty"`
	Amount float32 `json:"amount,omitempty"`
}
