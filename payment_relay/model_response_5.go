package payment_relay

type Response5 struct {
	// Unique Payment Reference Id
	PaymentRefId string `json:"payment_ref_id"`
	// Consumer Details
	Consumer []Response4Consumer `json:"consumer"`
	// Product Details
	Product []Response4Product `json:"product"`
	// Provider Details
	Provider []Response4Provider `json:"provider"`
	// User Details like hash id, email and any other user related info
	User []Response4User `json:"user"`
	Links []Response5Links `json:"links"`
}
