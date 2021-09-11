package core_payments

type RemittanceBasicFormV1Key struct {
	KeyName string `json:"key_name,omitempty"`
	FieldName string `json:"field_name,omitempty"`
	FieldKey string `json:"field_key,omitempty"`
	Type_ string `json:"type,omitempty"`
	Validation *RemittanceBasicFormV1KeyValidation `json:"validation,omitempty"`
	ValuesSupported []RemittanceBasicFormV1KeyValuesSupported `json:"values_supported,omitempty"`
}
