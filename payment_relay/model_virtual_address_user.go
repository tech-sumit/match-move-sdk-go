package payment_relay

// User Details
type VirtualAddressUser struct {
	// User MM Identifier
	HashId string `json:"hash_id,omitempty"`
	// User Virtual Payment Address
	Vpa string `json:"vpa,omitempty"`
	// User Email
	Email string `json:"email,omitempty"`
	// User Mobile Country Code
	MobileCountryCode string `json:"mobile_country_code,omitempty"`
	// User Mobile Number
	Mobile string `json:"mobile,omitempty"`
	// User Aadhaar ID
	Aadhaar string `json:"aadhaar,omitempty"`
	// Provider User Code
	Code string `json:"code,omitempty"`
}
