package core_payments

type AuthenticationsStatusesBody struct {
	// User Hash ID
	HashId string `json:"hash_id,omitempty"`
	// Approved or rejected status
	Status string `json:"status"`
	// Sub Status
	SubStatus string `json:"sub_status,omitempty"`
	// Message
	Message string `json:"message,omitempty"`
	// JSON encoded text string containing the KYC details
	KycDetails string `json:"kyc_details,omitempty"`
	// KYC Type
	KycType string `json:"kyc_type,omitempty"`
}
