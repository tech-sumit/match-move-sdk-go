package core_payments

type CurrencyV1 struct {
	// Currency ID
	Id string `json:"id,omitempty"`
	// Currency Name
	Name string `json:"name,omitempty"`
	// Currency Type
	Type_ string `json:"type,omitempty"`
	// Short name for the currency. For monetary currency use ISO-4217 
	Code string `json:"code,omitempty"`
	Date *CurrencyV1Date `json:"date,omitempty"`
}
