package core_payments

type FundCategoriesLimitResponseV1 struct {
	// Fund Category name
	Name string `json:"name,omitempty"`
	Settings *FundCategoriesLimitRequestV1Settings `json:"settings,omitempty"`
	Date *FundCategoriesLimitResponseV1Date `json:"date,omitempty"`
}
