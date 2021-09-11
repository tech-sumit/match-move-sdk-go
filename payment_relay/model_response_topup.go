package payment_relay

type ResponseTopup struct {
	// Currency code [ISO 3166](https://www.iso.org/iso-3166-country-codes.html) alpha-3
	Currency string `json:"currency,omitempty"`
	// Topup Amount
	Amount float32 `json:"amount,omitempty"`
}
