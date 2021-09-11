package payment_relay

type Response7 struct {
	// Name of payment provider
	Name string `json:"name"`
	// Product Code
	ProductCode string `json:"product_code,omitempty"`
	// Available payment types
	ProviderType string `json:"provider_type"`
	// Hash Value of Provider
	ProviderHash string `json:"provider_hash,omitempty"`
	// Provider Currency
	ProviderCurrency string `json:"provider_currency,omitempty"`
	// Mode of provider
	ProviderMode string `json:"provider_mode,omitempty"`
	// Date created
	DateCreated string `json:"date_created,omitempty"`
	// Topup minimum amount
	TopupMinAmount string `json:"topup_min_amount"`
	// Topup fee details
	FeeDetails []Response7FeeDetails `json:"fee_details"`
	// Description
	Description string `json:"description"`
	// Image
	Image []Response7Image `json:"image"`
}
