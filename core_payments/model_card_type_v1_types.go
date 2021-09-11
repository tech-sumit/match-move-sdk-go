package core_payments

type CardTypeV1Types struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	FormFactor string `json:"form_factor,omitempty"`
	PersonalisedOrder bool `json:"personalised_order,omitempty"`
	Token *CardFullV1Token `json:"token,omitempty"`
	Currency []string `json:"currency,omitempty"`
	Image *CardImageV1 `json:"image,omitempty"`
	Details *PaymentInstrumentDetailsV1 `json:"details,omitempty"`
}
