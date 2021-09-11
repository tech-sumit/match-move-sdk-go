package core_payments

type RemittanceCountriesV1 struct {
	// Array of countries supported in the program
	Countries []RemittanceCountriesV1Countries `json:"countries,omitempty"`
	Count *RemittanceCountriesV1Count `json:"count,omitempty"`
}
