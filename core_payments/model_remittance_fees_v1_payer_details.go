package core_payments

type RemittanceFeesV1PayerDetails struct {
	Code string `json:"code,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	PartnerCode string `json:"partner_code,omitempty"`
	PartnerAlphaCode string `json:"partner_alpha_code,omitempty"`
	PartnerDisplayName string `json:"partner_display_name,omitempty"`
	RoundingDecimalNumber string `json:"rounding_decimal_number,omitempty"`
	Version string `json:"version,omitempty"`
}
