package core_payments

// Date Limit settings
type FundCategoriesLimitResponseV1Date struct {
	// When was limit last updated
	Updated string `json:"updated,omitempty"`
	// When was limit last added
	Added string `json:"added,omitempty"`
}
