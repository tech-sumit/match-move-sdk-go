package payment_relay

type EnetsCallbackRequest struct {
	// Payment Reference ID
	MerchantTxnRef string `json:"merchant_txn_ref"`
	// Enets Reference ID
	NetsTxnRef string `json:"nets_txn_ref"`
	// Enets Status Code
	NetsTxnStatus string `json:"nets_txn_status"`
	// Enets Response Code
	NetsTxnRespCode string `json:"nets_txn_resp_code"`
	// Enets Response Message
	NetsTxnMsg string `json:"nets_txn_msg"`
	// Enets Recorded Transaction Time
	NetsTxnTime string `json:"nets_txn_time"`
	// Bank ID
	BankId string `json:"bank_id"`
	// Bank Reference ID
	BankRefCode string `json:"bank_ref_code"`
	// Bank Response Message
	BankRemarks string `json:"bank_remarks"`
	// Bank Status Code
	BankStatus string `json:"bank_status"`
	// Bank Recorded Transaction Date
	BankTxnDate string `json:"bank_txn_date"`
	// Bank Recorded Transaction Time
	BankTxnTime string `json:"bank_txn_time"`
	// Amount to topup
	TxnAmount string `json:"txnAmount"`
	// Source Exchange Rate
	SrcExchgRate string `json:"src_exchg_rate"`
	// Destination Exchange Rate
	DestExchgRate string `json:"dest_exchg_rate"`
	// Source Transaction Amount
	SrcTxnAmount string `json:"src_txn_amount"`
	// Source Transaction Currency Code
	SrcCurrencyCode string `json:"src_currency_code"`
	// Destination Transaction Currency Code
	DestCurrencyCode string `json:"dest_currency_code"`
	// JSON encoded text string containing additional parameters for confirmation hook
	AdditionalHook string `json:"additional_hook,omitempty"`
}
