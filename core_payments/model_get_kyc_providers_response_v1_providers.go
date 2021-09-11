package core_payments

type GetKycProvidersResponseV1Providers struct {
	// Provider Hash ID
	Id string `json:"id,omitempty"`
	// Provider Code
	Code string `json:"code,omitempty"`
	// Type of User (e.g customer, corporate, corporate_representative)
	UserType string `json:"user_type,omitempty"`
	// Callback URL
	CallbackUrl string `json:"callback_url,omitempty"`
	// Additional Details
	Details *interface{} `json:"details,omitempty"`
}
