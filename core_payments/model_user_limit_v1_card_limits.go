package core_payments

type UserLimitV1CardLimits struct {
	Interval string `json:"interval,omitempty"`
	Amount float64 `json:"amount,omitempty"`
	Count float64 `json:"count,omitempty"`
}
