package core_payments

type CardLimitsResponseV1CardTypeIssuanceLimitsPostKyc struct {
	// Card issuance is allowed or not
	Allowed bool `json:"allowed,omitempty"`
	// Number of cards can be created daily for post kyc user
	DailyCountLimit float64 `json:"daily_count_limit,omitempty"`
	// Number of cards can be created weekly for post kyc user
	WeeklyCountLimit float64 `json:"weekly_count_limit,omitempty"`
	// Number of cards can be created monthly for post kyc user
	MonthlyCountLimit float64 `json:"monthly_count_limit,omitempty"`
	// Number of cards can be created yearly for post kyc user
	YearlyCountLimit float64 `json:"yearly_count_limit,omitempty"`
	// Number of cards can be created lifetime for post kyc user
	LifetimeCountLimit float64 `json:"lifetime_count_limit,omitempty"`
}
