package core_payments

type BalanceCategoryV1CategoryName struct {
	Available *BalanceRootV1 `json:"available,omitempty"`
	Withholding *BalanceRootV1 `json:"withholding,omitempty"`
}
