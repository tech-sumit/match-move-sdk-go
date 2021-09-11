package core_payments

type InlineResponse20051Details struct {
	// Reference number given by the remittance provider.
	UniqueReferenceNumber string `json:"unique_reference_number,omitempty"`
	// Transaction code needed by the user to claim the remittance.
	TransactionCode string `json:"transaction_code,omitempty"`
	// Partner confirmation number needed by the user to claim the remittance (optional).
	PartnerConfirmationNumber string `json:"partner_confirmation_number,omitempty"`
}
