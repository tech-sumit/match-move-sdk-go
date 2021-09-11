package core_payments

type InlineResponse20055 struct {
	// For consumer payout agents
	SenderForm []RemittanceBeneficiaryFormV1 `json:"sender_form,omitempty"`
	// For corporate payout agents
	SendingBusiness []RemittanceBeneficiaryFormV1 `json:"sending_business,omitempty"`
	// For both consumer and corporate payout agents
	BeneficiaryForm []RemittanceBeneficiaryFormV1 `json:"beneficiary_form,omitempty"`
	// For both consumer and corporate payout agents
	AdditionalFields []RemittanceBeneficiaryFormV1 `json:"additional_fields,omitempty"`
}
