package core_payments

type PaymentInstrumentDetailsV1 struct {
	MinLoadLimit float64 `json:"min_load_limit,omitempty"`
	MaxLoadLimit float64 `json:"max_load_limit,omitempty"`
	NetworkType string `json:"network_type,omitempty"`
	Fee string `json:"fee,omitempty"`
	FeeCurrency string `json:"fee_currency,omitempty"`
	MaxCount string `json:"max_count,omitempty"`
	TopupLimits *PaymentInstrumentDetailsV1TopupLimits `json:"topup_limits,omitempty"`
	DeductLimits *PaymentInstrumentDetailsV1TopupLimits `json:"deduct_limits,omitempty"`
	TransferInLimits *PaymentInstrumentDetailsV1TopupLimits `json:"transfer_in_limits,omitempty"`
	TransferOutLimits *PaymentInstrumentDetailsV1TopupLimits `json:"transfer_out_limits,omitempty"`
	LoadLimits *PaymentInstrumentDetailsV1TopupLimits `json:"load_limits,omitempty"`
	UnloadLimits *PaymentInstrumentDetailsV1TopupLimits `json:"unload_limits,omitempty"`
	IssuanceLimits *CardIssuanceV1 `json:"issuance_limits,omitempty"`
	ActiveLimits *CardactiveLimitsV1 `json:"active_limits,omitempty"`
}
