package core_payments

type CardFullV1Funds struct {
	Available *BalanceRootV1 `json:"available,omitempty"`
	Withholding *BalanceRootV1 `json:"withholding,omitempty"`
	Categories *BalanceCategoryV1 `json:"categories,omitempty"`
}
