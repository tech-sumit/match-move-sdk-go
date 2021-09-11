package payment_relay

type HostedProviderHashBody struct {
	// Consumer Hash ID
	ConsumerHash string `json:"consumer_hash"`
	// OP Product Code
	ProductCode string `json:"product_code"`
	// Topup amount
	Amount float32 `json:"amount"`
	// User KYC Status [`pre_kyc`, `post_kyc`]
	KycStatus string `json:"kyc_status"`
	// User Unique Identifier
	UserHashId string `json:"user_hash_id"`
	// User Email Address (For Razorpay) Can be used as auto-populated contact email in their client-side form
	Email string `json:"email,omitempty"`
	// User Mobile Country Code (For Razorpay) Can be used as auto-populated contact number in their client-side form
	MobileCountryCode int32 `json:"mobile_country_code,omitempty"`
	// User Mobile Number (For Razorpay) Can be used as auto-populated contact number in their client-side form
	Mobile int32 `json:"mobile,omitempty"`
	// (For Razorpay) Can be used as auto-populated user name in their client-side form
	PreferredName string `json:"preferred_name,omitempty"`
	// (For Ingenico) (Required) Ingenico Product ID
	InRoute int32 `json:"in_route,omitempty"`
	// Currency code [ISO 3166](https://www.iso.org/iso-3166-country-codes.html) alpha-3
	Currency string `json:"currency,omitempty"`
	// (For Ingenico) [ISO 3166](https://www.iso.org/iso-3166-country-codes.html) alpha-2 country code
	CountryCode string `json:"country_code,omitempty"`
	// User's Card Primary Account Number
	Pan string `json:"pan,omitempty"`
	// Consumer Identifier for a Payment Transaction
	ClientRefId string `json:"client_ref_id,omitempty"`
	// Client IP Address
	ClientIp string `json:"client_ip,omitempty"`
	// Origin on where the payment transaction is executed [`admin`, `client`, `cron`, `payment_provider`]
	ActionOrigin string `json:"action_origin,omitempty"`
}
