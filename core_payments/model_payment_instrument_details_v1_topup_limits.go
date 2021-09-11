package core_payments

type PaymentInstrumentDetailsV1TopupLimits struct {
	PreKyc *PaymentLimitsV1 `json:"pre_kyc,omitempty"`
	PostKyc *PaymentLimitsV1 `json:"post_kyc,omitempty"`
	UserLimits *PaymentLimitsV1 `json:"user_limits,omitempty"`
}
