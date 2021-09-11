package core_payments

// A standard Wallet model
type WalletFullV1 struct {
	// Unique Reference for Wallet in the MatchMove System
	Id string `json:"id,omitempty"`
	// Account Identification Number for the Wallet
	Number string `json:"number,omitempty"`
	Funds *WalletFullV1Funds `json:"funds,omitempty"`
	Categories *BalanceCategoryV1 `json:"categories,omitempty"`
	Holder *CardFullV1Holder `json:"holder,omitempty"`
	Date *WalletFullV1Date `json:"date,omitempty"`
	Status *CardFullV1Status `json:"status,omitempty"`
	Image *WalletFullV1Image `json:"image,omitempty"`
	Details *PaymentInstrumentDetailsV1 `json:"details,omitempty"`
	Kit string `json:"kit,omitempty"`
}
