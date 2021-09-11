package core_payments

type BasicEnumerationIdentification struct {
	Code string `json:"code,omitempty"`
	Description *BasicEnumerationIdentificationDescription `json:"description,omitempty"`
}
