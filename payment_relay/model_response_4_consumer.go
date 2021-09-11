package payment_relay

type Response4Consumer struct {
	// Consumer Hash
	HashId string `json:"hash_id,omitempty"`
	// Consumer Name
	Name string `json:"name,omitempty"`
}
