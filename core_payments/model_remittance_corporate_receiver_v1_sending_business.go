package core_payments

type RemittanceCorporateReceiverV1SendingBusiness struct {
	Id string `json:"id,omitempty"`
	Ern string `json:"ern,omitempty"`
	RegisteredName string `json:"registered_name,omitempty"`
	TradingName string `json:"trading_name,omitempty"`
	Email string `json:"email,omitempty"`
	Msisdn string `json:"msisdn,omitempty"`
	RegistrationNumber string `json:"registration_number,omitempty"`
	TaxId string `json:"tax_id,omitempty"`
	Address *RemittanceCorporateReceiverV1SendingBusinessAddress `json:"address,omitempty"`
	RepresentativeFirstName string `json:"representative_first_name,omitempty"`
	RepresentativeMiddleName string `json:"representative_middle_name,omitempty"`
	RepresentativeLastName string `json:"representative_last_name,omitempty"`
	RepresentativeLastName2 string `json:"representative_last_name2,omitempty"`
	RepresentativeNativeName string `json:"representative_native_name,omitempty"`
	Identification *RemittanceCorporateReceiverV1SendingBusinessIdentification `json:"identification,omitempty"`
	Bank *RemittanceCorporateReceiverV1SendingBusinessBank `json:"bank,omitempty"`
	PayoutAgent string `json:"payout_agent,omitempty"`
	Product string `json:"product,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	IsActive float64 `json:"is_active,omitempty"`
	DateOfIncorporation string `json:"date_of_incorporation,omitempty"`
}
