package core_payments

type ErrorV1 struct {
	Code string `json:"code"`
	Description string `json:"description"`
	Link string `json:"link"`
}
