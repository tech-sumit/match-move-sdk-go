package core_payments

type TransactionRulesV1Inner struct {
	Id string `json:"id,omitempty"`
	Attribute string `json:"attribute,omitempty"`
	Action string `json:"action,omitempty"`
	Value string `json:"value,omitempty"`
	UserState string `json:"user_state,omitempty"`
	CardTypeCode string `json:"card_type_code,omitempty"`
}
