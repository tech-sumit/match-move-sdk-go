package payment_relay

type DirectProviderHashBody struct {
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
	// User Email Address
	Email string `json:"email,omitempty"`
	// User Mobile Country Code
	MobileCountryCode int32 `json:"mobile_country_code,omitempty"`
	// User Mobile Number
	Mobile int32 `json:"mobile,omitempty"`
	// Currency code ISO 4217 alpha code
	Currency string `json:"currency,omitempty"`
	// Token assigned to a particular card of the current user
	CardToken string `json:"card_token,omitempty"`
	// User's Card Primary Account Number
	Pan string `json:"pan,omitempty"`
	// Expiry Month of User's Card Primary Account Number
	ExpiryMonth string `json:"expiry_month,omitempty"`
	// Expiry Year of User's Card Primary Account Number
	ExpiryYear string `json:"expiry_year,omitempty"`
	// Security Code of User's Card Primary Account Number
	SecurityCode string `json:"security_code,omitempty"`
	// Source on where we received the transaction [`CALL_CENTRE`, `CARD_PRESENT`, `INTERNET`, `MAIL_ORDER`, `MOTO`, `TELEPHONE_ORDER`, `VOICE_RESPONSE`]
	Source string `json:"source,omitempty"`
	// JSON-encoded value of User's billing address which should contain [street, street2, city, stateProvince, country(3 letter ISO standard alpha country code), postcodeZip]
	BillingAddress string `json:"billing_address,omitempty"`
	// JSON-encoded value of User's Details which should contain [firstName, lastName]
	Customer string `json:"customer,omitempty"`
	// Consumer Identifier for a Payment Transaction
	ClientRefId string `json:"client_ref_id,omitempty"`
	// Client IP Address
	ClientIp string `json:"client_ip,omitempty"`
	// Origin on where the payment transaction is executed [`admin`, `client`, `cron`, `payment_provider`]
	ActionOrigin string `json:"action_origin,omitempty"`
	// Will save card
	SaveCard string `json:"save_card,omitempty"`
}
