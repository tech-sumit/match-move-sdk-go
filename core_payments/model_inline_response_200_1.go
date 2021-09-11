package core_payments

type InlineResponse2001 struct {
	Transactions []AnyOfinlineResponse2001TransactionsItems `json:"transactions,omitempty"`
	Links []TransactionLinksV1 `json:"links,omitempty"`
	// Records per page returned. Echoed back from the request param
	RecordsPerPage int32 `json:"records_per_page,omitempty"`
	// Page number of the transactions returned
	Page int32 `json:"page,omitempty"`
	// Total number of records for the query params
	TotalRecords int32 `json:"total_records,omitempty"`
}
