package core_payments

type RemittanceCountriesV1Currency struct {
	// ISO-4217 three-letter currency code
	Code string `json:"code,omitempty"`
	// ISO-4217 numeric currency code
	Numeric string `json:"numeric,omitempty"`
}
