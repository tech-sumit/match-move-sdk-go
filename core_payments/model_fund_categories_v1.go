package core_payments

type FundCategoriesV1 struct {
	// Unique ID of the fund category in the system
	Id string `json:"id"`
	// Assigned name for the fund category
	Name string `json:"name"`
	// The max allowed transaction on the fund category
	DefaultAmount string `json:"default_amount"`
	Factor string `json:"factor"`
	Mcc string `json:"mcc,omitempty"`
	Merchant string `json:"merchant,omitempty"`
	Terminal string `json:"terminal,omitempty"`
}
