package core_payments

type BalanceCategoryV1 struct {
	CATEGORY_NAME *BalanceCategoryV1CategoryName `json:"CATEGORY_NAME,omitempty"`
	DEFAULT *BalanceCategoryV1CategoryName `json:"DEFAULT"`
}
