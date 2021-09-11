package payment_relay

// Topup Details
type Response3Details struct {
	Reference *Response3DetailsReference `json:"reference,omitempty"`
	Topup *Response3DetailsTopup `json:"topup,omitempty"`
	Fees *Response3DetailsFees `json:"fees,omitempty"`
	User *Response3DetailsUser `json:"user,omitempty"`
}
