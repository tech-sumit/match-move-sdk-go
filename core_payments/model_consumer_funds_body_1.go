package core_payments

type ConsumerFundsBody1 struct {
	// Reference identifiers for the debit transaction(s). Multiple identifiers in CSV format are accepted.
	Ids string `json:"ids"`
}
