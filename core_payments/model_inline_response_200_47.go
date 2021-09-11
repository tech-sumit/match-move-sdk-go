package core_payments

type InlineResponse20047 struct {
	TotalPages float64 `json:"total_pages,omitempty"`
	RecordsPerPage float64 `json:"records_per_page,omitempty"`
	Page string `json:"page,omitempty"`
	TotalRecords float64 `json:"total_records,omitempty"`
	Links *InlineResponse20047Links `json:"links,omitempty"`
	Data []BankAccountV1 `json:"data,omitempty"`
}
