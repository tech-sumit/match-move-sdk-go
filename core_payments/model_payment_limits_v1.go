package core_payments

type PaymentLimitsV1 struct {
	Allowed float64 `json:"allowed,omitempty"`
	LifetimeCountLimit float64 `json:"lifetime_count_limit,omitempty"`
	LifetimeTransactionalLimit float64 `json:"lifetime_transactional_limit,omitempty"`
	DailyCountLimit float64 `json:"daily_count_limit,omitempty"`
	DailyTransactionalLimit float64 `json:"daily_transactional_limit,omitempty"`
	WeeklyCountLimit float64 `json:"weekly_count_limit,omitempty"`
	WeeklyTransactionalLimit float64 `json:"weekly_transactional_limit,omitempty"`
	MonthlyCountLimit float64 `json:"monthly_count_limit,omitempty"`
	MonthlyTransactionalLimit float64 `json:"monthly_transactional_limit,omitempty"`
	UnitTransactionLimit float64 `json:"unit_transaction_limit,omitempty"`
}
