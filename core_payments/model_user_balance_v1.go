package core_payments

type UserBalanceV1 struct {
	Wallet *UserBalanceV1Wallet `json:"wallet,omitempty"`
	Card *UserBalanceV1Card `json:"card,omitempty"`
}
