package core_payments

type InlineResponse20049 struct {
	TotalPages float64 `json:"total_pages,omitempty"`
	RecordsPerPage float64 `json:"records_per_page,omitempty"`
	Page string `json:"page,omitempty"`
	TotalRecords float64 `json:"total_records,omitempty"`
	Links *InlineResponse20049Links `json:"links,omitempty"`
	Data []FundTransferV1 `json:"data,omitempty"`
}
