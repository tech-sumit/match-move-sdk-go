package core_payments

type CardFullV1Type struct {
	// One of the card types available via GET /users/wallets/cards/types
	Type_ string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	FormFactor string `json:"form_factor,omitempty"`
	PersonalisedOrder bool `json:"personalised_order,omitempty"`
}
