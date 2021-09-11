package core_payments

type InlineResponse20019 struct {
	Transactions []AnyOfinlineResponse20019TransactionsItems `json:"transactions,omitempty"`
	Links []TransactionLinksV1 `json:"links,omitempty"`
	Count *InlineResponse20019Count `json:"count,omitempty"`
}
