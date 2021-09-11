package core_payments

type RemittanceBeneficiaryFormV1 struct {
	FieldGroupName string `json:"field_group_name,omitempty"`
	FieldName string `json:"field_name,omitempty"`
	FieldKey string `json:"field_key,omitempty"`
	Type_ string `json:"type,omitempty"`
	Validation *RemittanceBeneficiaryFormV1Validation `json:"validation,omitempty"`
	ValuesSupported []RemittanceBasicFormV1KeyValuesSupported `json:"values_supported,omitempty"`
}
