package payment_relay

type TransactionWhitelist struct {
	// Fees Hash Id
	Id string `json:"id"`
	// Consumer Name
	ConsumerName string `json:"consumer_name"`
	// Product Code
	ProductCode string `json:"product_code"`
	// Provider Name
	ProviderName string `json:"provider_name"`
	// User Hash ID
	UserHashId string `json:"user_hash_id"`
	// Topup amount
	Amount float64 `json:"amount"`
	// Name entered from payment provider screen, if any
	CustName string `json:"cust_name"`
	// Source type from payment screen whether entered Card, UPI, etc
	SourceType string `json:"source_type"`
	// Source type issuer name
	SourceName string `json:"source_name"`
	// Source account number or address
	SourceNumber string `json:"source_number"`
	// Source account expiry, if applicable
	SourceExpiry string `json:"source_expiry"`
	// Source brand
	SourceBrand string `json:"source_brand"`
	// Source funding method
	SourceFundingMethod string `json:"source_funding_method"`
	// Is active
	IsActive bool `json:"is_active"`
	// Date created
	DateCreated string `json:"date_created"`
	// Date modified
	DateModified string `json:"date_modified"`
}
