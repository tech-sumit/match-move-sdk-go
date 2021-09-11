package payment_relay

type InlineResponse200Provider struct {
	// Consumer Hash
	HashId string `json:"hash_id,omitempty"`
	// Provider Name
	Name string `json:"name,omitempty"`
	Channel *InlineResponse200Channel `json:"channel,omitempty"`
}
