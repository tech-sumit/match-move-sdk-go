package core_payments

type UserReportV1 struct {
	// Users information
	Users []AnyOfUserReportV1UsersItems `json:"users,omitempty"`
	TotalRecords string `json:"total_records,omitempty"`
	RecordsPerPage string `json:"records_per_page,omitempty"`
	TotalPages string `json:"total_pages,omitempty"`
	Links []ClosedLoopV1Links `json:"links,omitempty"`
}
