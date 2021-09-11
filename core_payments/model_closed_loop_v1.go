package core_payments

type ClosedLoopV1 struct {
	TotalRecords int32 `json:"total_records,omitempty"`
	RecordsPerPage int32 `json:"records_per_page,omitempty"`
	TotalPages int32 `json:"total_pages,omitempty"`
	Links []ClosedLoopV1Links `json:"links,omitempty"`
	// Transaction information
	Transactions []ClosedLoopTransactionsV1 `json:"transactions,omitempty"`
}
