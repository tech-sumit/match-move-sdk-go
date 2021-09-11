package payment_relay

type Refunds struct {
	// Refund ID
	Id int32 `json:"id"`
	// Topup/Credit ID
	CreditId int32 `json:"credit_id"`
	// Payment reference ID
	PaymentRefId string `json:"payment_ref_id"`
	// Refund reference ID
	RefundRefId string `json:"refund_ref_id"`
	// Provider reference ID
	ProviderRefId string `json:"provider_ref_id"`
	// Refund type (full, partial)
	Type_ string `json:"type"`
	// During which action refund transaction is created [topup, stale, blacklist, cron, settlement]
	Action string `json:"action"`
	// Refund amount
	Amount int32 `json:"amount"`
	// reason for refund
	Reason string `json:"reason"`
	// Number of refund attempts
	Attempts int32 `json:"attempts"`
	// Status [`0`, `1`, `2`, `3`, `4`, `5`, `6`]
	Status int32 `json:"status"`
	// Date created `(YYYY-MM-DD HH:MM:SS)`
	DateCreated string `json:"date_created"`
	// Date modified `(YYYY-MM-DD HH:MM:SS)`
	DateModified string `json:"date_modified"`
}
