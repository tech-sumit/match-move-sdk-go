package core_payments

type ConsumerFundsBody struct {
	// Reference identifiers for the credit transaction(s). Multiple identifiers in CSV format are accepted.
	Ids string `json:"ids"`
}
