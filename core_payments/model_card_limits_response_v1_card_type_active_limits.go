package core_payments

type CardLimitsResponseV1CardTypeActiveLimits struct {
	// Card issuance is allowed or not
	Allowed bool `json:"allowed,omitempty"`
	// Number of active cards for pre kyc user
	PreKyc float64 `json:"pre_kyc,omitempty"`
	Statuses []string `json:"statuses,omitempty"`
	// Number of active cards for post kyc user
	PostKyc float64 `json:"post_kyc,omitempty"`
}
