package core_payments

type InlineResponse20024 struct {
	Transfers []TransactionFundsTransferV1 `json:"transfers"`
	Count *InlineResponse20024Count `json:"count"`
}
