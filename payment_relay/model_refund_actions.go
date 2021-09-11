package payment_relay

type RefundActions struct {
	// Refund action ID
	Id int32 `json:"id"`
	// Refund ID
	RefundId int32 `json:"refund_id"`
	// The action e.g. processing, enquiring etc.
	ActionResult string `json:"action_result"`
	// Origin of the action (admin, cron, client, payment_provider, topup)
	ActionOrigin string `json:"action_origin"`
	// General human readable description of the action
	ActionMessage string `json:"action_message"`
	// Remarks
	Remarks string `json:"remarks"`
	// Date created `(YYYY-MM-DD HH:MM:SS)`
	DateCreated string `json:"date_created"`
	// Date modified `(YYYY-MM-DD HH:MM:SS)`
	DateModified string `json:"date_modified"`
}
