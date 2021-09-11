package core_payments

type CardactiveLimitsV1 struct {
	Allowed float64 `json:"allowed,omitempty"`
	PreKyc float64 `json:"pre_kyc,omitempty"`
	PostKyc float64 `json:"post_kyc,omitempty"`
	Statuses []string `json:"statuses,omitempty"`
}
