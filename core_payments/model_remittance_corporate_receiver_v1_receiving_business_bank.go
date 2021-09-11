package core_payments

type RemittanceCorporateReceiverV1ReceivingBusinessBank struct {
	Account *RemittanceCorporateReceiverV1SendingBusinessBankAccount `json:"account,omitempty"`
	Branch *RemittanceCorporateReceiverV1SendingBusinessBankBranch `json:"branch,omitempty"`
}
