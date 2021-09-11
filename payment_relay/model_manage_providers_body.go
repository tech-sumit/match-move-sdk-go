package payment_relay

type ManageProvidersBody struct {
	// provider name
	ProviderName string `json:"provider_name"`
	// Provider type
	Type_ string `json:"type"`
	// Provider mode
	ProviderMode string `json:"provider_mode"`
	// Currency code [ISO 3166](https://www.iso.org/iso-3166-country-codes.html) alpha-3
	Currency string `json:"currency"`
	// Hashing Algorithm [PHP](http://php.net/manual/en/function.hash-algos.php)
	Algorithm string `json:"algorithm,omitempty"`
	// Provider description
	Description string `json:"description"`
}
