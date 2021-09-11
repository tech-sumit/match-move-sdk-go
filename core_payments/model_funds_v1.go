package core_payments

type FundsV1 struct {
	// Hash ID of the transaction
	Id string `json:"id,omitempty"`
	// Reference id of the transaction
	RefId string `json:"ref_id,omitempty"`
	// Status of the transaction
	Status string `json:"status,omitempty"`
}
