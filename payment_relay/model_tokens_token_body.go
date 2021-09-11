package payment_relay

type TokensTokenBody struct {
	// Callback Signature for security
	Signature string `json:"signature"`
	// Topup amount
	Amount float32 `json:"amount"`
	// Currency code [ISO 3166](https://www.iso.org/iso-3166-country-codes.html) alpha-3
	Currency string `json:"currency"`
	// Consumer Identifier for a Payment Transaction
	ClientRefId string `json:"client_ref_id,omitempty"`
	// JSON encoded text string containing additional parameters for confirmation hook
	AdditionalHook string `json:"additional_hook,omitempty"`
}
