package core_payments

type CardLimitsResponseV1CardType struct {
	IssuanceLimits *CardLimitsResponseV1CardTypeIssuanceLimits `json:"issuance_limits,omitempty"`
	ActiveLimits *CardLimitsResponseV1CardTypeActiveLimits `json:"active_limits,omitempty"`
}
