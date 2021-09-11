package core_payments

type OauthTransactionV1 struct {
	Field string `json:"field,omitempty"`
	Type_ string `json:"type,omitempty"`
	SubType string `json:"sub_type,omitempty"`
	Details string `json:"details,omitempty"`
	OldRemittanceFloatingBalance string `json:"old_remittance_floating_balance,omitempty"`
	NewRemittanceFloatingBalance string `json:"new_remittance_floating_balance,omitempty"`
	OldFundFloatingBalance string `json:"old_fund_floating_balance,omitempty"`
	NewFundFloatingBalance string `json:"new_fund_floating_balance,omitempty"`
	OldFundDebitLimit string `json:"old_fund_debit_limit,omitempty"`
	NewFundDebitLimit string `json:"new_fund_debit_limit,omitempty"`
	TransactionRefId string `json:"transaction_ref_id,omitempty"`
	Method string `json:"method,omitempty"`
	Uri string `json:"uri,omitempty"`
	DateAdded string `json:"date_added,omitempty"`
}
