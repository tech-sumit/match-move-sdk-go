package core_payments

type RemittanceCountriesV1Countries struct {
	// Country Name
	Name string `json:"name,omitempty"`
	// ISO-3166 alpha-3 country code
	Code string `json:"code,omitempty"`
	Currency *RemittanceCountriesV1Currency `json:"currency,omitempty"`
	Agents []interface{} `json:"agents,omitempty"`
}
