package core_payments

type RemittancePaymentModesV1 struct {
	// Array of payment channel/modes supported in the system
	Modes []RemittancePaymentModesV1Modes `json:"modes,omitempty"`
	Count *RemittancePaymentModesV1Count `json:"count,omitempty"`
}
