package core_payments

type CardsCardTypeCodeBody struct {
	// Required value supported only for non personalized physical cards. Physical Card Proxy Number [format: PY + Proxy Number (eg. 'PY000000000123')] or Account identification number. When present, this will bind the card to the current user. 
	AssocNumber string `json:"assoc_number,omitempty"`
	// Transaction reference identifier obtained after using `POST /users/wallets/funds` when the consumer is pre-funded or after using `POST /oauth/consumer/funds` when the consumer is not pre-funded. 
	RefId string `json:"ref_id,omitempty"`
	CardDesign string `json:"card_design,omitempty"`
	NameOnCard string `json:"name_on_card,omitempty"`
	AdditionalDetails string `json:"additional_details,omitempty"`
	// If passed as true the POST request does the activation of the card along with the linkage of the card. Do note this parameter is only supported for non personalized card. Also depending on the card management partner used only option `true` might be supported 
	AutoActivate bool `json:"auto_activate,omitempty"`
	Var2faMethod string `json:"2fa_method,omitempty"`
	Var2faDelivery string `json:"2fa_delivery,omitempty"`
	Var2faValue string `json:"2fa_value,omitempty"`
}
