package core_payments

type CardMiniV1 struct {
	// Unique Reference for Wallet in the MatchMove System
	Id string `json:"id"`
	Status *CardFullV1Status `json:"status"`
	Holder *CardFullV1Holder `json:"holder"`
	Funds *CardFullV1Funds `json:"funds,omitempty"`
	// Account Identification Number for the Wallet
	Number string `json:"number"`
	// Network the card is associated to Example : visa, mastercard, rupay
	Network string `json:"network,omitempty"`
	// One of the card types available via GET /users/wallets/cards/types
	CardTypeCode string `json:"card_type_code,omitempty"`
	Image *CardImageV1 `json:"image,omitempty"`
	Date *CardFullV1Date `json:"date,omitempty"`
}
