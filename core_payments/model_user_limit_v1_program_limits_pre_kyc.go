package core_payments

type UserLimitV1ProgramLimitsPreKyc struct {
	Amount float64 `json:"amount,omitempty"`
	Count float64 `json:"count,omitempty"`
	Interval string `json:"interval,omitempty"`
}
