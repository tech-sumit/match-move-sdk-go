package core_payments

type RemittanceCorporateReceiverV1SendingBusinessBank struct {
	Account *RemittanceCorporateReceiverV1SendingBusinessBankAccount `json:"account,omitempty"`
	Branch *RemittanceCorporateReceiverV1SendingBusinessBankBranch `json:"branch,omitempty"`
}
