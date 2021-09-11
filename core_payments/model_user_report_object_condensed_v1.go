package core_payments

type UserReportObjectCondensedV1 struct {
	// User Identifier
	Id string `json:"id,omitempty"`
	// Email Address
	Email string `json:"email,omitempty"`
	Name *UserReportObjectCondensedV1Name `json:"name,omitempty"`
	Mobile *UserReportObjectCondensedV1Mobile `json:"mobile,omitempty"`
	Identification *UserReportObjectCondensedV1Identification `json:"identification,omitempty"`
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
	TotalBalances *UserBalanceV1 `json:"total_balances,omitempty"`
	Links []CardFullV1Links `json:"links,omitempty"`
}
