package core_payments

type CardsCardIdBody2 struct {
	// Values can be `lost`, `stolen` and `damaged`. These values will permanently block the card and is irreversible. If not specified , it will temporarily block the card.
	Type_ string `json:"type,omitempty"`
}
