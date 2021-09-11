package core_payments

type TransactionCategoriesRulesResponseV1 struct {
	// transcation category hash id from transcation_categories
	Id string `json:"id,omitempty"`
	// transcation category is enable or disbale
	Enable bool `json:"enable,omitempty"`
	Rules *TransactionCategoriesRulesRequestV1 `json:"rules,omitempty"`
}
