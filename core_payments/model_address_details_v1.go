package core_payments

type AddressDetailsV1 struct {
	// Address line 2 for the address type (Do note please sanitize any trailing or leading spaces before posting the data)
	Address1 string `json:"address_1"`
	// Address line 2 for the address type (Do note please sanitize any trailing or leading spaces before posting the data)
	Address2 string `json:"address_2"`
	// Address line 2 for the address type (Do note please sanitize any trailing or leading spaces before posting the data)
	Address3 string `json:"address_3,omitempty"`
	// Address line 2 for the address type (Do note please sanitize any trailing or leading spaces before posting the data)
	Address4 string `json:"address_4,omitempty"`
	// City for the address type
	City string `json:"city"`
	// State for the address type
	State string `json:"state"`
	// Country. *Case-sensitive* value referenced as `country_name` which can be obtained from mysql-country-list
	Country string `json:"country"`
	// Country code in ISO 3166 alpha-3 format. (Examples : SGP , PHL, IND) 
	CountryCode string `json:"country_code,omitempty"`
	// Zip Code for the address type
	Zipcode string `json:"zipcode"`
}
