package core_payments

type RemittanceCorporateReceiverV1SendingBusinessAddress struct {
	Address1 string `json:"address_1,omitempty"`
	Address2 string `json:"address_2,omitempty"`
	City string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
	Country *AmountDetailsV1Country `json:"country,omitempty"`
	Zipcode string `json:"zipcode,omitempty"`
}
