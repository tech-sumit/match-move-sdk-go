package payment_relay

type ConsumerHashProviderHashBody struct {
	// OP Product Code
	ProductCode string `json:"product_code"`
	// Will inactivate consumer rate
	Inactivate string `json:"inactivate"`
}
