package payment_relay

type Response4User struct {
	// User Hash
	HashId string `json:"hash_id,omitempty"`
	// User Email Address
	Email string `json:"email,omitempty"`
}
