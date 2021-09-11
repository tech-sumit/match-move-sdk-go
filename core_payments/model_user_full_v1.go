package core_payments

type UserFullV1 struct {
	Id string `json:"id"`
	// Registration email for the user to be created in the MatchMove system for a particular program.
	Email string `json:"email"`
	Name *UserNameV1 `json:"name"`
	Mobile *UserMobileV1 `json:"mobile"`
	Identification *UserIdentificationV1 `json:"identification,omitempty"`
	// Country of issue of registration documents for the users in the MatchMove system
	Countryofissue string `json:"countryofissue,omitempty"`
	// Country of issue of registration documents for the users in the MatchMove system
	CountryCodeOfIssue string `json:"country_code_of_issue,omitempty"`
	// Birthday of the user registered in the MatchMove system. The value is expected to match the value provided in documents for the KYC process.
	Birthday string `json:"birthday,omitempty"`
	PlaceOfBirth string `json:"place_of_birth,omitempty"`
	Gender string `json:"gender,omitempty"`
	Nationality string `json:"nationality,omitempty"`
	Title string `json:"title,omitempty"`
	MaritalStatus string `json:"marital_status,omitempty"`
	MothersMaidenName string `json:"mothers_maiden_name,omitempty"`
	CustomerId string `json:"customer_id,omitempty"`
	Status *CardFullV1Status `json:"status"`
	Date *UserFullV1Date `json:"date"`
	PartnerId string `json:"partner_id,omitempty"`
}
