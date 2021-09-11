package payment_relay

type MmaDpayCallbackRequest struct {
	// Timestamp of Request
	RESPONSE_DATE_TIME string `json:"RESPONSE_DATE_TIME"`
	// Response Code
	RESPONSE_CODE string `json:"RESPONSE_CODE"`
	// Customer User Hash ID
	CUST_ID string `json:"CUST_ID"`
	// MMADpay Authorization Code
	AUTH_CODE string `json:"AUTH_CODE,omitempty"`
	// Customer Mobile Number
	CUST_PHONE string `json:"CUST_PHONE"`
	// MMADpay MOP Type
	MOP_TYPE string `json:"MOP_TYPE,omitempty"`
	// Masked Card Number
	CARD_MASK string `json:"CARD_MASK,omitempty"`
	// Currency Code in ISO 4217 numeric format
	CURRENCY_CODE string `json:"CURRENCY_CODE"`
	// Bank Reference Number
	RRN string `json:"RRN,omitempty"`
	// MMADpay Transaction Message
	PG_TXN_MESSAGE string `json:"PG_TXN_MESSAGE,omitempty"`
	// Transaction Status
	STATUS string `json:"STATUS"`
	// Product Description
	PRODUCT_DESC string `json:"PRODUCT_DESC"`
	// Unique Identifier by IRCTC iPay
	PG_REF_NUM string `json:"PG_REF_NUM,omitempty"`
	// Amount
	AMOUNT string `json:"AMOUNT"`
	// Response Message
	RESPONSE_MESSAGE string `json:"RESPONSE_MESSAGE"`
	// Customer Email
	CUST_EMAIL string `json:"CUST_EMAIL"`
	// Unique Transaction Identifier by IRCTC iPay
	TXN_ID string `json:"TXN_ID"`
	// MMADpay Acq ID
	ACQ_ID string `json:"ACQ_ID,omitempty"`
	// Transaction Type
	TXNTYPE string `json:"TXNTYPE"`
	// Flag whether surcharge is applied
	SURCHARGE_FLAG string `json:"SURCHARGE_FLAG,omitempty"`
	// Security Hash for each unique transaction
	HASH string `json:"HASH"`
	// Payment type
	PAYMENT_TYPE string `json:"PAYMENT_TYPE,omitempty"`
	// Callback URL
	RETURN_URL string `json:"RETURN_URL"`
	// Merchant ID
	PAY_ID string `json:"PAY_ID"`
	// Payment Reference ID
	ORDER_ID string `json:"ORDER_ID"`
	// Original Unique Transaction Identifier by IRCTC iPay
	ORIG_TXN_ID string `json:"ORIG_TXN_ID,omitempty"`
	// MMADpay AVR
	AVR string `json:"AVR,omitempty"`
	// Customer Name
	CUST_NAME string `json:"CUST_NAME"`
}
