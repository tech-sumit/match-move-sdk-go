package core_payments

type UserBalanceV1WalletAvailable struct {
	Currency string `json:"currency,omitempty"`
	Amount float32 `json:"amount,omitempty"`
}
