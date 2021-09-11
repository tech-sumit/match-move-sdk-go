package payment_relay

type VirtualAddress struct {
	// Virtual Address Hash Id
	Id string `json:"id"`
	// Unique Identifier from Provider
	TransactionId string `json:"transaction_id"`
	Consumer *FeesConsumer `json:"consumer"`
	Provider *FeesProvider `json:"provider"`
	User *VirtualAddressUser `json:"user"`
	Bank *VirtualAddressBank `json:"bank"`
	// Is active
	IsActive bool `json:"is_active"`
	// Date created
	DateCreated string `json:"date_created"`
	// Date modified
	DateModified string `json:"date_modified"`
}
