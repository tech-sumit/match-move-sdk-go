package core_payments

type InlineResponse20012 struct {
	Transactions []AnyOfinlineResponse20012TransactionsItems `json:"transactions,omitempty"`
	Links []TransactionLinksV1 `json:"links,omitempty"`
	// Records per page returned. Echoed back from the request param
	RecordsPerPage string `json:"records_per_page,omitempty"`
	// Page number of the transactions returned
	Page string `json:"page,omitempty"`
	// Total number of records for the query params
	TotalRecords string `json:"total_records,omitempty"`
}
