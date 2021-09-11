package payment_relay

// Refund Details
type TransactionItemRefundDetails struct {
	// MM Settlement Refund Status
	Status string `json:"status,omitempty"`
	// MM Settlement Refund Status Message
	Message string `json:"message,omitempty"`
	// Original Credit Status from Provider
	OriginalPgStatus string `json:"original_pg_status,omitempty"`
	// Original Credit Status from MM
	OriginalMmStatus string `json:"original_mm_status,omitempty"`
	// Refund Status from Provider
	SettlementPgStatus string `json:"settlement_pg_status,omitempty"`
}
