package core_payments

type InlineResponse20040 struct {
	Id string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	Name *UserNameV1 `json:"name,omitempty"`
	Mobile *UserMobileV1 `json:"mobile,omitempty"`
	Links []interface{} `json:"links,omitempty"`
	Status *CardFullV1Status `json:"status,omitempty"`
	Date *InlineResponse20040Date `json:"date,omitempty"`
}
