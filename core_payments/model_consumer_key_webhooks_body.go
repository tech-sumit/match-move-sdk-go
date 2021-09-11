package core_payments

type ConsumerKeyWebhooksBody struct {
	Url string `json:"url,omitempty"`
	EventHash string `json:"event_hash,omitempty"`
	CategoryHash string `json:"category_hash,omitempty"`
}
