package core_payments

type RefundTypeBody struct {
	Amount int32 `json:"amount,omitempty"`
	Details string `json:"details,omitempty"`
}
