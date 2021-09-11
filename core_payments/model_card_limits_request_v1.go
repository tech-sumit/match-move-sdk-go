package core_payments

type CardLimitsRequestV1 struct {
	// Card statuses
	ActiveCountStatuses string `json:"active_count_statuses"`
	AppendStatus bool `json:"append_status,omitempty"`
	// User kyc state
	UserState string `json:"user_state"`
	// Number of cards can be created daily
	DailyCountLimit string `json:"daily_count_limit"`
	// Number of cards can be created weekly
	WeeklyCountLimit string `json:"weekly_count_limit"`
	// Number of cards can be created monthly
	MonthlyCountLimit string `json:"monthly_count_limit"`
	// Number of cards can be created yearly
	YearlyCountLimit string `json:"yearly_count_limit"`
	// Number of cards can be created lifetime
	LifetimeCountLimit string `json:"lifetime_count_limit"`
	// Number of cards can be active at a time
	ActiveCountLimit string `json:"active_count_limit"`
	// Meta data for setting the rules for card creation
	IssuanceCountRules string `json:"issuance_count_rules,omitempty"`
}
