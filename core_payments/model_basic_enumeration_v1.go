package core_payments

type BasicEnumerationV1 struct {
	// The code for the enumeration type to be used
	Code string `json:"code"`
	// The description for the specific enumeration.
	Description string `json:"description,omitempty"`
}
