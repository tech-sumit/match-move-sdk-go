package payment_relay

// Payment Source details
type TransactionStatusDetailsSource struct {
	// Customer MM User Hash ID
	CustId string `json:"cust_id,omitempty"`
	// Email entered from payment provider screen, if any
	CustEmail string `json:"cust_email,omitempty"`
	// Mobile entered from payment provider screen, if any
	CustPhone string `json:"cust_phone,omitempty"`
	// Name entered from payment provider screen, if any
	CustName string `json:"cust_name,omitempty"`
	// Payment screen url before provider initiated callback to RELAY
	Referrer string `json:"referrer,omitempty"`
	// Source type from payment screen whether entered Card, UPI, etc
	Type_ string `json:"type,omitempty"`
	// Source type issuer name
	Name string `json:"name,omitempty"`
	// Source account number or address
	AccountNumber string `json:"account_number,omitempty"`
	// Source account expiry, if applicable
	Expiry string `json:"expiry,omitempty"`
	// Source brand
	Brand string `json:"brand,omitempty"`
	// Source funding method
	FundingMethod string `json:"funding_method,omitempty"`
	// Source authentication
	Authentication string `json:"authentication,omitempty"`
}
