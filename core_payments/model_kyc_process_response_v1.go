package core_payments

type KycProcessResponseV1 struct {
	// Indicator if the process was successful
	Success bool `json:"success"`
	Links []KycProcessResponseV1Links `json:"links"`
}
