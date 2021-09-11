package core_payments

type RemittanceAgentV1 struct {
	// Payout Agent Hash ID
	AgentId string `json:"agent_id"`
	PaymentMode *RemittanceAgentV1PaymentMode `json:"payment_mode"`
	// Payout Agent Display Name
	DisplayName string `json:"display_name"`
	// Provider Rails
	Provider string `json:"provider"`
	// Calculate exchange rate after fx percentage discounting
	ExchangeRate string `json:"exchange_rate"`
	// ISO-3166 alpha-3 country code
	CountryCode string `json:"country_code"`
	// ISO-4217 three-letter receiving currency code
	ReceiveCurrency string `json:"receive_currency"`
	// Receiver/Beneficiary Type
	ReceivingAgentType string `json:"receiving_agent_type"`
	ExchangesRate *RemittanceAgentV1ExchangesRate `json:"exchanges_rate"`
	Fees *RemittanceAgentV1Fees `json:"fees"`
	AgentDetails *RemittanceAgentV1AgentDetails `json:"agent_details"`
	Links []CardFullV1Links `json:"links,omitempty"`
}
