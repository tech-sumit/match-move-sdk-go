package core_payments

type UserAddressV1 struct {
	Residential *AddressDetailsV1 `json:"residential,omitempty"`
	Billing *AddressDetailsV1 `json:"billing,omitempty"`
}
