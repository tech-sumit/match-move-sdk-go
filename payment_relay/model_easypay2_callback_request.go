package payment_relay

type Easypay2CallbackRequest struct {
	// Merchant ID assigned by Easypay2
	TMMCode string `json:"TM_MCode"`
	// Payment Reference ID -- The value in this field is an echo of the `ref` field submitted from IH-Relay to Easypay2
	TMRefNo string `json:"TM_RefNo"`
	// Amount -- The value in this field is an echo of the `amt` field submitted from IH-Relay to Easypay2 (In json format, this must not be enclosed in quotation marks serving as string)
	TMDebitAmt float32 `json:"TM_DebitAmt"`
	// Currency -- The value in this field is an echo of the `cur` field submitted from IH-Relay to Easypay2
	TMCurrency string `json:"TM_Currency"`
	// Status of the Transaction : Valid values         `NO` – Unsuccessful (Bank rejected transaction), `YES` – Successful (Bank has authorized the authorization / payment)
	TMStatus string `json:"TM_Status"`
	// Bank’s approval code for successful transactions NULL or empty if the transaction failed.
	TMApprovalCode string `json:"TM_ApprovalCode"`
	// Response code returned by the acquiring bank         `00` for successful transactions, `NN` denotes failure reason code
	TMBankRespCode string `json:"TM_BankRespCode"`
	// Error details
	TMError string `json:"TM_Error,omitempty"`
	// Additional user field 1
	TMUserField1 string `json:"TM_UserField1,omitempty"`
	// Additional user field 2
	TMUserField2 string `json:"TM_UserField2,omitempty"`
	// Additional user field 3
	TMUserField3 string `json:"TM_UserField3,omitempty"`
	// Additional user field 4
	TMUserField4 string `json:"TM_UserField4,omitempty"`
	// Additional user field 5 : Callback Signature
	TMUserField5 string `json:"TM_UserField5"`
	// Last 4 digits of credit card number
	TMCCLast4Digit string `json:"TM_CCLast4Digit"`
	// Credit card expiry date specified in request
	TMExpiryDate string `json:"TM_ExpiryDate"`
	// Type of transaction performed
	TMTrnType string `json:"TM_TrnType"`
	// Sub-type of transaction performed
	TMSubTrnType string `json:"TM_SubTrnType,omitempty"`
	// This value is returned if rcard is specified from IH-Relay to Easypay2
	TMCardNum string `json:"TM_CardNum"`
	// Transaction pay type
	TMOriginalPayType string `json:"TM_OriginalPayType,omitempty"`
	// JSON encoded text string containing additional parameters for confirmation hook
	AdditionalHook string `json:"additional_hook,omitempty"`
}
