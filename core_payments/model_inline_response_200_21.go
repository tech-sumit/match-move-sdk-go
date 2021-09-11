package core_payments

type InlineResponse20021 struct {
	// Transaction reference identifier
	Id string `json:"id"`
	// Transaction reference identifier
	RefId string `json:"ref_id"`
	// User id of the user for whcih the transaction was performed
	UserId string `json:"user_id"`
	// Wallet identifier
	WalletId string `json:"wallet_id"`
	// Card identifier
	CardId string `json:"card_id"`
	// Indicator if the transaction needs to be confirmed or not. Values can be [require or skip]
	Confirm string `json:"confirm"`
}
