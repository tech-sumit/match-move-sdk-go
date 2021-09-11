package core_payments

type CardFullV1 struct {
	// Unique Reference for Wallet in the MatchMove System
	Id string `json:"id"`
	// Account Identification Number for the Wallet
	Number string `json:"number"`
	Holder *CardFullV1Holder `json:"holder"`
	Type_ *CardFullV1Type `json:"type,omitempty"`
	Token *CardFullV1Token `json:"token,omitempty"`
	Date *CardFullV1Date `json:"date,omitempty"`
	Image *CardImageV1 `json:"image,omitempty"`
	Status *CardFullV1Status `json:"status"`
	Funds *CardFullV1Funds `json:"funds,omitempty"`
	Links []CardFullV1Links `json:"links,omitempty"`
	Kit string `json:"kit,omitempty"`
}
