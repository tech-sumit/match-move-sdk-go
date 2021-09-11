package core_payments

type CardLimitsResponseV1CardTypeIssuanceLimitsPreKyc struct {
	// Card issuance is allowed or not
	Allowed bool `json:"allowed,omitempty"`
	// Number of cards can be created daily for pre kyc user
	DailyCountLimit float64 `json:"daily_count_limit,omitempty"`
	// Number of cards can be created weekly for pre kyc user
	WeeklyCountLimit float64 `json:"weekly_count_limit,omitempty"`
	// Number of cards can be created monthly for pre kyc user
	MonthlyCountLimit float64 `json:"monthly_count_limit,omitempty"`
	// Number of cards can be created yearly for pre kyc user
	YearlyCountLimit float64 `json:"yearly_count_limit,omitempty"`
	// Number of cards can be created lifetime for pre kyc user
	LifetimeCountLimit float64 `json:"lifetime_count_limit,omitempty"`
}
