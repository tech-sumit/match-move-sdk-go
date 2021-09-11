package payment_relay

type Activities struct {
	// Record ID
	Id int32 `json:"id"`
	// Consumer ID
	ConsumerId int32 `json:"consumer_id"`
	// IP address of user
	Ip string `json:"ip"`
	// Request details
	Details string `json:"details"`
	// Date created `(YYYY-MM-DD HH:MM:SS)`
	DateCreated string `json:"date_created"`
}
