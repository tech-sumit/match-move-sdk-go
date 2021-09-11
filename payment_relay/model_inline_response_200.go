package payment_relay

type InlineResponse200 struct {
	// Unique Payment Reference Id
	PaymentRefId string `json:"payment_ref_id,omitempty"`
	// Consumer Details
	Consumer []InlineResponse200Consumer `json:"consumer"`
	// Product Details
	Product []InlineResponse200Product `json:"product"`
	// Provider Details
	Provider []InlineResponse200Provider `json:"provider"`
	// User Details like hash id, email and any other user related info
	User []InlineResponse200User `json:"user"`
	// Fees breakdown
	Fees []InlineResponse200Fees `json:"fees"`
	// Expiry date of the transaction
	ExpiresAt string `json:"expires_at"`
	// Instuction or message from Relay to customer
	Message string `json:"message"`
	// Expiry duration in seconds
	ExpiryDuration int32 `json:"expiry_duration,omitempty"`
	QrDetails []InlineResponse200QrDetails `json:"qr_details,omitempty"`
	QrImage []InlineResponse200QrImage `json:"qr_image,omitempty"`
}
