package core_payments

type InlineResponse20022 struct {
	// The user ID.
	Id string `json:"id,omitempty"`
	RefId string `json:"ref_id,omitempty"`
	UserId string `json:"user_id,omitempty"`
	WalletId string `json:"wallet_id,omitempty"`
	Confirm string `json:"confirm,omitempty"`
	Links []interface{} `json:"links,omitempty"`
}
