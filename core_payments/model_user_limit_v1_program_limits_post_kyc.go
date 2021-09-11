package core_payments

type UserLimitV1ProgramLimitsPostKyc struct {
	Amount float64 `json:"amount,omitempty"`
	Interval string `json:"interval,omitempty"`
	Count float64 `json:"count,omitempty"`
}
