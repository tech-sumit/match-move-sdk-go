package core_payments

type InlineResponse20043 struct {
	TotalPages float64 `json:"total_pages,omitempty"`
	RecordsPerPage float64 `json:"records_per_page,omitempty"`
	Page float64 `json:"page,omitempty"`
	TotalRecords float64 `json:"total_records,omitempty"`
	Links *InlineResponse20043Links `json:"links,omitempty"`
	Data []VirtualAccountV1 `json:"data,omitempty"`
}
