package core_payments

type InlineResponse20014 struct {
	// Unique Reference for Wallet in the MatchMove System
	Id string `json:"id,omitempty"`
	// Account Identification Number for the Wallet
	Number string `json:"number,omitempty"`
	Funds *CardFullV1Funds `json:"funds,omitempty"`
	Holder *CardFullV1Holder `json:"holder,omitempty"`
	Date *InlineResponse20014Date `json:"date,omitempty"`
	Status *CardFullV1Status `json:"status,omitempty"`
	Cards []CardMiniV1 `json:"cards,omitempty"`
}
