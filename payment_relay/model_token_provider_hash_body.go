package payment_relay

type TokenProviderHashBody struct {
	// Consumer Hash ID
	ConsumerHash string `json:"consumer_hash"`
	// OP Product Code
	ProductCode string `json:"product_code"`
	// Provider Channel
	ProviderChannel string `json:"provider_channel,omitempty"`
	// Topup amount
	Amount float64 `json:"amount"`
	// User KYC Status [pre_kyc, post_kyc]
	KycStatus string `json:"kyc_status"`
	// User Unique Identifier
	UserHashId string `json:"user_hash_id"`
	// User Email Address
	Email string `json:"email,omitempty"`
	// User Mobile Country Code
	MobileCountryCode int32 `json:"mobile_country_code,omitempty"`
	// User Mobile Number
	Mobile int32 `json:"mobile,omitempty"`
	// Currency code ISO 3166 alpha-3
	Currency string `json:"currency,omitempty"`
	// User’s Card Primary Account Number
	Pan string `json:"pan,omitempty"`
	// Consumer Identifier for a Payment Transaction
	ClientRefId string `json:"client_ref_id,omitempty"`
	// Origin on where the payment transaction is executed [admin, client, cron, payment_provider]
	ActionOrigin string `json:"action_origin,omitempty"`
}
