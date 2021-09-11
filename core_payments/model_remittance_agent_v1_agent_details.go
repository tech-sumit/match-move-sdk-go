package core_payments

type RemittanceAgentV1AgentDetails struct {
	// Payout Agent Hash ID
	AgentId string `json:"agent_id"`
	// Calculation Mode - either source or receive
	CalculationMode string `json:"calculation_mode"`
	// Assigned payout agent channel code
	PaymentMode string `json:"payment_mode"`
	// Assigned payout agent routing param
	RoutingParam string `json:"routing_param"`
	// Assigned payout agent routing type
	RoutingType string `json:"routing_type"`
	// Supported rails currency
	ProviderCurrency string `json:"provider_currency"`
	// ISO-4217 three-letter sending currency code
	SendCurrency string `json:"send_currency"`
	// ISO-4217 three-letter receiving currency code
	ReceiveCurrency string `json:"receive_currency"`
	// ISO-3166 alpha-3 receiving country code
	ReceiveCountry string `json:"receive_country"`
	// Payout agent partner name - usually the bank name
	PartnerName string `json:"partner_name"`
	// Assigned priority
	Priority string `json:"priority,omitempty"`
	// Minimum amount allowed for remittance transaction
	MinAmount string `json:"min_amount,omitempty"`
	// Maximum amount allowed for remittance transaction
	MaxAmount string `json:"max_amount,omitempty"`
}
