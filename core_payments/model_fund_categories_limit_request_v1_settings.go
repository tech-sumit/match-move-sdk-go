package core_payments

// Limit settings
type FundCategoriesLimitRequestV1Settings struct {
	// Duration of the limit
	Duration string `json:"duration,omitempty"`
	// Interval of the limit
	Interval string `json:"interval,omitempty"`
	// Value of the limit
	Value string `json:"value,omitempty"`
}
