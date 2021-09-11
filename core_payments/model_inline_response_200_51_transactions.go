package core_payments

type InlineResponse20051Transactions struct {
	// Reference ID of the transaction that was confirmed.
	Id string `json:"id,omitempty"`
	// Status of the confirmation.
	Status string `json:"status,omitempty"`
	// Describes the result of the confirmation.
	Description string `json:"description,omitempty"`
	Details []InlineResponse20051Details `json:"details,omitempty"`
}
