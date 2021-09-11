package core_payments

type RemittanceFeesV1 struct {
	Fees []RemittanceFeesV1Fees `json:"fees,omitempty"`
}
