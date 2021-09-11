package core_payments

type UserRequestV1 struct {
	Id string `json:"id,omitempty"`
	// First name or given name
	FirstName string `json:"first_name,omitempty"`
	// Last name or family name
	LastName string `json:"last_name,omitempty"`
	// User Middle name
	MiddleName string `json:"middle_name,omitempty"`
	// Prefereed name
	PreferredName string `json:"preferred_name,omitempty"`
	// Registration email for the user to be created in the MatchMove system for a particular program.
	Email string `json:"email,omitempty"`
	// Supported Country codes for the program can be retrieved from `GET /users/enumerations/mobile_country_codes` 
	MobileCountryCode string `json:"mobile_country_code,omitempty"`
	// Minimum 8 Characters , One special character , One number , One small character 
	Password string `json:"password,omitempty"`
	// Mobile numbe rof the user
	Mobile string `json:"mobile,omitempty"`
	// Title to be set like mr/miss etc
	Title string `json:"title,omitempty"`
	// Id type to be set. Available values can be retrieved from `GET /users/enumerations/id_types`
	IdType string `json:"id_type,omitempty"`
	// Id number to be unique
	IdNumber string `json:"id_number,omitempty"`
	// Id type expiry date
	IdDateExpiry string `json:"id_date_expiry,omitempty"`
	// Issued date of id type
	IdDateIssue string `json:"id_date_issue,omitempty"`
	// Country that issued your id type
	CountryOfIssue string `json:"country_of_issue,omitempty"`
	// alternative id type
	AltIdType string `json:"alt_id_type,omitempty"`
	// Alertnative id type number to be unique
	AltIdNumber string `json:"alt_id_number,omitempty"`
	// Alternative id type expiry date
	AltIdDateExpiry string `json:"alt_id_date_expiry,omitempty"`
	// Alternative id type issued date
	AltIdDateIssue string `json:"alt_id_date_issue,omitempty"`
	// Alternative Id type isunace country
	AltIdCountryOfIssue string `json:"alt_id_country_of_issue,omitempty"`
	// Birthday of the user registred in matchmove system
	Birthday string `json:"birthday,omitempty"`
	// Place where you have born
	PlaceOfBirth string `json:"place_of_birth,omitempty"`
	// User Nationality. Supported values can be retrieved from `GET /users/enumerations/nationalities`
	Nationality string `json:"nationality,omitempty"`
	// User Gender. Supported values can be retreived from `GET /users/enumerations/genders`
	Gender string `json:"gender,omitempty"`
	// User marital Status. Supported values can be retreived from `GET /users/enumerations/marital_status`
	MaritalStatus string `json:"marital_status,omitempty"`
	// Mother's maiden name
	MothersMaidenName string `json:"mothers_maiden_name,omitempty"`
	// Nature of your business
	NatureOfBusiness string `json:"nature_of_business,omitempty"`
	// Customer id to be passed
	CustomerId string `json:"customer_id,omitempty"`
	// Purpose of your account
	PurposeOfAccount string `json:"purpose_of_account,omitempty"`
	// Partner id to be passed
	PartnerId string `json:"partner_id,omitempty"`
}
