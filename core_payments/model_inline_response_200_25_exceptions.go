package core_payments

type InlineResponse20025Exceptions struct {
	Id string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Type_ string `json:"type"`
	DateAdded string `json:"date_added"`
	DateExpiry string `json:"date_expiry"`
	Links *InlineResponse20025Links `json:"links,omitempty"`
}
