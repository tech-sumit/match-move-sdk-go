package core_payments

type RemittanceAgentV1ExchangesRate struct {
	// Exchange Rate value from rails
	Value string `json:"value"`
	// Exchange Rate expiry
	Expiry string `json:"expiry,omitempty"`
}
