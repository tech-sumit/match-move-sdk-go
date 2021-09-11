package core_payments

type KycProcessRequestV1 struct {
	// Provider ID from the selected provider for the KYC flow
	ProviderId string `json:"provider_id"`
	// Type of user enabled for the selected provider
	UserType string `json:"user_type"`
	// OP User Hash ID
	UserId string `json:"user_id"`
	// OP Program Code
	ProgramCode string `json:"program_code"`
}
