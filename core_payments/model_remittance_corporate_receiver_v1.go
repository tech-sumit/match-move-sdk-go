package core_payments

type RemittanceCorporateReceiverV1 struct {
	BeneficiaryHashId string `json:"beneficiary_hash_id,omitempty"`
	SendingBusiness *RemittanceCorporateReceiverV1SendingBusiness `json:"sending_business,omitempty"`
	ReceivingBusiness *RemittanceCorporateReceiverV1ReceivingBusiness `json:"receiving_business,omitempty"`
}
