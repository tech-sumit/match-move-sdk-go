package core_payments

type BasicEnumerationIdentificationDescription struct {
	ProcessorRegexValidation string `json:"processor_regex_validation,omitempty"`
	MaxDuplicate float64 `json:"max_duplicate,omitempty"`
	RegexValidation string `json:"regex_validation,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Masking bool `json:"masking,omitempty"`
}
