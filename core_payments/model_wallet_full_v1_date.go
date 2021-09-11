package core_payments

type WalletFullV1Date struct {
	// Expiry of the payment instrument
	Expiry string `json:"expiry,omitempty"`
	// Date of issue of payment instrument
	Issued string `json:"issued,omitempty"`
	Closed string `json:"closed,omitempty"`
}
