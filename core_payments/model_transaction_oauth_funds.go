package core_payments

type TransactionOauthFunds struct {
	Id string `json:"id,omitempty"`
	RefId string `json:"ref_id,omitempty"`
	Status string `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
	SorTransaction *TransactionOauthFundsSorTransaction `json:"sor_transaction,omitempty"`
}
