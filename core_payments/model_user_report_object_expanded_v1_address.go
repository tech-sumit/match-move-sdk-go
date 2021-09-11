package core_payments

type UserReportObjectExpandedV1Address struct {
	Residential *UserAddressV1 `json:"residential,omitempty"`
	Billing *UserAddressV1 `json:"billing,omitempty"`
}
