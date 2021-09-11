package core_payments

type UserLimitV1 struct {
	// Auto generated hash id
	Id string `json:"id,omitempty"`
	// Transaction is enable or not for that card
	Enable bool `json:"enable,omitempty"`
	// Whether User has set any limits for that card
	UserDetailsSet bool `json:"user_details_set,omitempty"`
	// Cards per interval limits value
	CardLimits []UserLimitV1CardLimits `json:"card_limits,omitempty"`
	ProgramLimits *UserLimitV1ProgramLimits `json:"program_limits,omitempty"`
}
