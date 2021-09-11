package payment_relay

type InlineResponse200User struct {
	// User Hash
	HashId string `json:"hash_id,omitempty"`
	// User Email Address
	Email string `json:"email,omitempty"`
}
