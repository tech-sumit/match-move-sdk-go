package core_payments

// Number of cards can be issued to a particular user
type CardLimitsResponseV1CardTypeIssuanceLimits struct {
	PreKyc *CardLimitsResponseV1CardTypeIssuanceLimitsPreKyc `json:"pre_kyc,omitempty"`
	PostKyc *CardLimitsResponseV1CardTypeIssuanceLimitsPostKyc `json:"post_kyc,omitempty"`
	Rules []string `json:"rules,omitempty"`
}
