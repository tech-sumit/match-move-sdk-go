package core_payments

type TransactionsRulesBody struct {
	Attribute string `json:"attribute,omitempty"`
	Action string `json:"action,omitempty"`
	Value string `json:"value,omitempty"`
	CardTypeCode string `json:"card_type_code,omitempty"`
	UserState string `json:"user_state,omitempty"`
}
