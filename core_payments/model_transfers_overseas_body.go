package core_payments

type TransfersOverseasBody struct {
	// Reference identifiers for the overseas transfer transaction(s). Multiple identifiers in CSV format are accepted.
	Ids string `json:"ids,omitempty"`
}
