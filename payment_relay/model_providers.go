package payment_relay

type Providers struct {
	Razorpay *RazorpayCallbackRequest `json:"razorpay,omitempty"`
	Easypay2 *Easypay2CallbackRequest `json:"easypay2,omitempty"`
	Enets *EnetsCallbackRequest `json:"enets,omitempty"`
}
