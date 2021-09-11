package payment_relay

// User details
type TransactionStatusDetailsUser struct {
	// User Hash ID
	Id string `json:"id,omitempty"`
	// User Email
	Email string `json:"email,omitempty"`
	// User Mobile Country Code
	MobileCountryCode string `json:"mobile_country_code,omitempty"`
	// User Mobile
	Mobile string `json:"mobile,omitempty"`
	// User Name
	Name string `json:"name,omitempty"`
	// User OP Product Code
	ProductCode string `json:"product_code,omitempty"`
}
