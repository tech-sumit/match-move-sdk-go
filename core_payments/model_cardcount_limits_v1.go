package core_payments

type CardcountLimitsV1 struct {
	// whether its allowed or not
	Allowed float64 `json:"allowed,omitempty"`
	// Cards allowed to be created on daily basis
	DailyCountLimit float64 `json:"daily_count_limit,omitempty"`
	// Cards allowed to be created on weekly basis
	WeeklyCountLimit float64 `json:"weekly_count_limit,omitempty"`
	// Cards allowed to be created on monthly basis
	MonthlyCountLimit float64 `json:"monthly_count_limit,omitempty"`
	// Cards allowed to be created on yearly basis
	YearlyCountLimit float64 `json:"yearly_count_limit,omitempty"`
	// Cards allowed to be created for lifetime of account
	LifetimeCountLimit float64 `json:"lifetime_count_limit,omitempty"`
}
