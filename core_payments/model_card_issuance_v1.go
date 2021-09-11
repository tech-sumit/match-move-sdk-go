package core_payments

type CardIssuanceV1 struct {
	PreKyc *CardcountLimitsV1 `json:"pre_kyc,omitempty"`
	PostKyc *CardcountLimitsV1 `json:"post_kyc,omitempty"`
	Rules []string `json:"rules,omitempty"`
}
