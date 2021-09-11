package core_payments

type UserReportObjectCondensedV1Identification struct {
	IdType string `json:"id_type,omitempty"`
	IdNumber string `json:"id_number,omitempty"`
	IssueDate string `json:"issue_date,omitempty"`
	ExpiryDate string `json:"expiry_date,omitempty"`
}
