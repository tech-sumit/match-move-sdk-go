package payment_relay

// Consumer details of who modified last
type FeesModifiedBy struct {
	// Consumer hash
	ConsumerHash string `json:"consumer_hash,omitempty"`
	// Consumer name
	ConsumerName string `json:"consumer_name,omitempty"`
}
