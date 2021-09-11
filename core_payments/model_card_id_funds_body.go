package core_payments

type CardIdFundsBody struct {
	Amount int32 `json:"amount"`
	Message string `json:"message,omitempty"`
	FundCategoryName string `json:"fund_category_name,omitempty"`
}
