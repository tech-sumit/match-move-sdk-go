package core_payments

type RemittanceDetailsV1Sender struct {
	FirstName string `json:"first_name,omitempty"`
	MiddleName string `json:"middle_name,omitempty"`
	LastName string `json:"last_name,omitempty"`
	Mobile string `json:"mobile,omitempty"`
}
