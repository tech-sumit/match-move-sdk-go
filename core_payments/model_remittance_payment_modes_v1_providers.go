package core_payments

type RemittancePaymentModesV1Providers struct {
	// Provider code
	Name string `json:"name,omitempty"`
	// Channel code
	Code string `json:"code,omitempty"`
	// Provider channel code
	ProviderCode string `json:"provider_code,omitempty"`
	// Array of countries which supports the given payment mode
	Countries []interface{} `json:"countries,omitempty"`
}
