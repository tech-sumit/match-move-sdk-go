package core_payments

type InlineResponse200 struct {
	// The user id on which the actions was performed
	Id string `json:"id,omitempty"`
	State string `json:"state,omitempty"`
}
