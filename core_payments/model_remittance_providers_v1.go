package core_payments

type RemittanceProvidersV1 struct {
	Types []RemittanceProvidersV1Types `json:"types,omitempty"`
	Count *RemittanceProvidersV1Count `json:"count,omitempty"`
}
