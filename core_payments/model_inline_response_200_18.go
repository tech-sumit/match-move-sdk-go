package core_payments

type InlineResponse20018 struct {
	Success string `json:"success,omitempty"`
	SubStatus string `json:"sub_status,omitempty"`
	KycDetails string `json:"kyc_details,omitempty"`
	KycType string `json:"kyc_type,omitempty"`
	Message string `json:"message,omitempty"`
}
