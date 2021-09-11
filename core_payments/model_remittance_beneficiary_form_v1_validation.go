package core_payments

type RemittanceBeneficiaryFormV1Validation struct {
	MaxLength int32 `json:"max_length,omitempty"`
	IsRequired string `json:"is_required,omitempty"`
}
