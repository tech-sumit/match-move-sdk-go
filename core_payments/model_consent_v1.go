package core_payments

type ConsentV1 struct {
	FullName string `json:"full_name,omitempty"`
	Numbers string `json:"numbers,omitempty"`
	Sign string `json:"sign,omitempty"`
}
