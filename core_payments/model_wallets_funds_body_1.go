package core_payments

type WalletsFundsBody1 struct {
	// The registered `email` address of the user for which the fund transaction is being initiated
	Email string `json:"email,omitempty"`
	// The `mobile_country_code` of the user for which the fund transaction is being initiated
	MobileCountryCode string `json:"mobile_country_code,omitempty"`
	// The `mobile` of the user for which the fund transaction is being initiated
	Mobile string `json:"mobile,omitempty"`
	// The `hash_id` of the user's wallet for which the fund transaction is being initiated
	HashId string `json:"hash_id,omitempty"`
	Amount int32 `json:"amount"`
	Details string `json:"details,omitempty"`
	FundCategoryName string `json:"fund_category_name,omitempty"`
}
