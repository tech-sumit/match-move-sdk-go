package core_payments

type RemittanceSendV1 struct {
	Status string `json:"status,omitempty"`
	RefId string `json:"ref_id,omitempty"`
	TransferDetails *RemittanceSendV1TransferDetails `json:"transfer_details,omitempty"`
	Quotation *RemittanceQuotationV1 `json:"quotation,omitempty"`
	DateAdded string `json:"date_added,omitempty"`
	DateLastUpdated string `json:"date_last_updated,omitempty"`
	Confirm string `json:"confirm,omitempty"`
	Links []TransactionLinksV1 `json:"links,omitempty"`
}
