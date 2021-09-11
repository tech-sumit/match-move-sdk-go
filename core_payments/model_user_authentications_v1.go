package core_payments

type UserAuthenticationsV1 struct {
	Documents []UserAuthenticationsV1Documents `json:"documents,omitempty"`
	Count float64 `json:"count,omitempty"`
	Status string `json:"status,omitempty"`
	SubStatus string `json:"sub_status,omitempty"`
	KycDetails string `json:"kyc_details,omitempty"`
}
