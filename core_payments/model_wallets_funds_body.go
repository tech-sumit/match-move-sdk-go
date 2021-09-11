package core_payments

type WalletsFundsBody struct {
	// The registered `email` address of the user for which the fund transaction is being initiated
	Email string `json:"email,omitempty"`
	// The `mobile_country_code` of the user for which the fund transaction is being initiated
	MobileCountryCode string `json:"mobile_country_code,omitempty"`
	// The `mobile` of the user for which the fund transaction is being initiated
	Mobile string `json:"mobile,omitempty"`
	// The `user_id` of the user for which the fund transaction is being initiated
	UserId string `json:"user_id,omitempty"`
	// The amount for the transaction. The format supported would be based on the precision of the base currency
	Amount float64 `json:"amount"`
	// JSON encoded details to be passed as meta data to be stored in MatchMove sytem
	Details string `json:"details,omitempty"`
	ForwardedFor string `json:"forwarded_for,omitempty"`
	HashedPan string `json:"hashed_pan,omitempty"`
	// Optional category to specify where the credit/transfer of funds happens.
	FundCategoryName string `json:"fund_category_name,omitempty"`
	Type_ string `json:"type,omitempty"`
	// To be specified if the funds need to be directly moved to the card in non-shared card balance.
	LoadCard bool `json:"load_card,omitempty"`
	// The card_id of the card in non-shared balance to which the funds are to be moved. Required if the value for `load_card` is true.
	CardId string `json:"card_id,omitempty"`
}
