package core_payments

type TransactionLinksV1 struct {
	Rel string `json:"rel,omitempty"`
	Href string `json:"href,omitempty"`
	Method string `json:"method,omitempty"`
}
