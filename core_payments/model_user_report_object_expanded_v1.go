package core_payments

type UserReportObjectExpandedV1 struct {
	// User Identifier
	Id string `json:"id,omitempty"`
	// Email Address
	Email string `json:"email,omitempty"`
	Name *UserNameV1 `json:"name,omitempty"`
	Mobile *UserMobileV1 `json:"mobile,omitempty"`
	Identification *UserIdentificationV1 `json:"identification,omitempty"`
	// Country of Issue
	Countryofissue string `json:"countryofissue,omitempty"`
	// Date of Birth
	Birthday string `json:"birthday,omitempty"`
	// Gender
	Gender string `json:"gender,omitempty"`
	// Title
	Title string `json:"title,omitempty"`
	CustomerId string `json:"customer_id,omitempty"`
	Status *UserReportObjectCondensedV1Status `json:"status,omitempty"`
	// Registration Date
	DateAdded string `json:"date_added,omitempty"`
	Address *UserReportObjectExpandedV1Address `json:"address,omitempty"`
	Documents *UserAuthenticationsV1 `json:"documents,omitempty"`
	TotalBalances *UserBalanceV1 `json:"total_balances,omitempty"`
	Links []CardFullV1Links `json:"links,omitempty"`
}
