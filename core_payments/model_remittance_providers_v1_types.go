package core_payments

type RemittanceProvidersV1Types struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Details string `json:"details,omitempty"`
	Links []CardFullV1Links `json:"links,omitempty"`
}
