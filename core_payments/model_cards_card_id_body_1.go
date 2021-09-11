package core_payments

type CardsCardIdBody1 struct {
	// Card unique identifier. When present, this will unblock the card associated with this parameter.
	Id string `json:"id"`
	// Physical Card Proxy Number [format: PY + Proxy Number (eg. 'PY000000000123')] or Account identification number. When present, this will bind the card to the current user. 
	AssocNumber string `json:"assoc_number,omitempty"`
	// Transaction reference identifier obtained after using POST /users/wallets/funds when the consumer is pre-funded or after using POST /oauth/consumer/funds when the consumer is not pre-funded.
	RefId string `json:"ref_id,omitempty"`
}
