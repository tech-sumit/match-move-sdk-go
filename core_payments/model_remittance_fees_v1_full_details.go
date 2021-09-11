package core_payments

type RemittanceFeesV1FullDetails struct {
	HashId string `json:"hash_id,omitempty"`
	Address string `json:"address,omitempty"`
	City string `json:"city,omitempty"`
	// Payment mode of the the agent
	CountryCode string `json:"country_code,omitempty"`
	// Payment mode of the the agent
	PaymentModeDescription string `json:"payment_mode_description,omitempty"`
	PayerCode string `json:"payer_code,omitempty"`
	MainTimeTable string `json:"main_time_table,omitempty"`
	SecondaryTimeTable string `json:"secondary_time_table,omitempty"`
	CityCode string `json:"city_code,omitempty"`
	Telephone1 string `json:"telephone_1,omitempty"`
	Telephone2 string `json:"telephone_2,omitempty"`
	Fax string `json:"fax,omitempty"`
	State string `json:"state,omitempty"`
	PaymentCurrencies string `json:"payment_currencies,omitempty"`
	Version string `json:"version,omitempty"`
	RequiredFields string `json:"required_fields,omitempty"`
	RequiredDocuments string `json:"required_documents,omitempty"`
	// The type of transactions the agent is supposed to receive
	BeneficiaryType string `json:"beneficiary_type,omitempty"`
	AgentName string `json:"agent_name,omitempty"`
}
