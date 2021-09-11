package core_payments

type InlineResponse20044 struct {
	TotalRecords int32 `json:"total_records,omitempty"`
	RecordsPerPage int32 `json:"records_per_page,omitempty"`
	TotalPages int32 `json:"total_pages,omitempty"`
	Page int32 `json:"page,omitempty"`
	Links []TransactionLinksV1 `json:"links,omitempty"`
	Data []VirtualAccountTransaction `json:"data,omitempty"`
}
