package core_payments

type BalanceRootV1 struct {
	// currency of the balance type in [ISO 4217] format (http://www.iso.org/iso/currency_codes)
	Currency string `json:"currency,omitempty"`
	// Monetary value of the balance type in the defined application currency with the relevant precision
	Amount string `json:"amount"`
	Name string `json:"name,omitempty"`
}
