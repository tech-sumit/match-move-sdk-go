package payment_relay

// User information
type Response3DetailsUser struct {
	// User email address
	Email string `json:"email,omitempty"`
	// User mobile country code
	MobileCountryCode int32 `json:"mobile_country_code,omitempty"`
	// User mobile number
	Mobile int32 `json:"mobile,omitempty"`
}
