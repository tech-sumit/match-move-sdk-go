package payment_relay

type ResponseDetails struct {
	Reference []ResponseReference `json:"reference,omitempty"`
	Topup []ResponseTopup `json:"topup,omitempty"`
	User []ResponseUser `json:"user,omitempty"`
}
