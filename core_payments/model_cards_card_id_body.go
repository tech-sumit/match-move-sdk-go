package core_payments

type CardsCardIdBody struct {
	// Card activation code obtained from response of POST /users/wallets/cards/{card_type.code}. 
	ActivationCode string `json:"activation_code,omitempty"`
	Var2faMethod string `json:"2fa_method,omitempty"`
	Var2faDelivery string `json:"2fa_delivery,omitempty"`
	Var2faValue string `json:"2fa_value,omitempty"`
}
