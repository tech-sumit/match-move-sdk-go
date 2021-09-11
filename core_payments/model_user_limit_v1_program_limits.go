package core_payments

// Cards per interval limits value
type UserLimitV1ProgramLimits struct {
	PreKyc []UserLimitV1ProgramLimitsPreKyc `json:"pre_kyc,omitempty"`
	PostKyc []UserLimitV1ProgramLimitsPostKyc `json:"post_kyc,omitempty"`
}
