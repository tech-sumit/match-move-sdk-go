package payment_relay

type Response8Items struct {
	Consumer *FeesConsumer `json:"consumer,omitempty"`
	Product *FeesProduct `json:"product,omitempty"`
	Provider *FeesProvider `json:"provider,omitempty"`
	ProviderChannel *FeesProviderChannel `json:"provider_channel,omitempty"`
	Fees *Response8Fees `json:"fees,omitempty"`
}
