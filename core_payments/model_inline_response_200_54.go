package core_payments

type InlineResponse20054 struct {
	Agents []RemittanceAgentV1 `json:"agents,omitempty"`
	Count *InlineResponse20054Count `json:"count,omitempty"`
	Link []InlineResponse20054Link `json:"link,omitempty"`
}
