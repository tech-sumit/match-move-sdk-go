package core_payments

type RemittanceCorporateReceiverV1SendingBusinessIdentification struct {
	Number string `json:"number,omitempty"`
	Type_ string `json:"type,omitempty"`
	CountryOfIssue string `json:"country_of_issue,omitempty"`
	IssueDate string `json:"issue_date,omitempty"`
	ExpireDate string `json:"expire_date,omitempty"`
}
