package core_payments

type UserMobileV1 struct {
	// Mobile Country code of the user
	CountryCode string `json:"country_code,omitempty"`
	// Mobile number of the user excluding the country code
	Number string `json:"number,omitempty"`
}
