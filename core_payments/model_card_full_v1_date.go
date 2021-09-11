package core_payments

type CardFullV1Date struct {
	Expiry string `json:"expiry,omitempty"`
	Issued string `json:"issued,omitempty"`
	Closed string `json:"closed,omitempty"`
}
