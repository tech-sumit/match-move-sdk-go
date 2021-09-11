package core_payments

type UserLimitRequestV1 struct {
	// We can disbale tarnsaction category for that card
	Enable bool `json:"enable,omitempty"`
	// Limits and number of count we can set per interval
	Values []UserLimitV1CardLimits `json:"values,omitempty"`
}
