package payment_relay

type RazorpayCallbackRequest struct {
	// Amount
	Amt string `json:"amt,omitempty"`
	// Consumer hash
	ConsumerHash string `json:"consumer_hash,omitempty"`
	// Product code
	ProductCode string `json:"product_code,omitempty"`
	// User email
	Email string `json:"email,omitempty"`
	// Payment Reference ID
	PaymentRefId string `json:"payment_ref_id"`
	// Razorpay Payment ID
	RazorpayPaymentId string `json:"razorpay_payment_id"`
	// Signature
	Signature string `json:"signature"`
	// JSON encoded text string containing additional parameters for confirmation hook
	AdditionalHook string `json:"additional_hook,omitempty"`
}
