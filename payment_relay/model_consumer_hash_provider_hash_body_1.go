package payment_relay

type ConsumerHashProviderHashBody1 struct {
	// product code
	ProductCode string `json:"product_code"`
	// fee flat
	FeeFlat string `json:"fee_flat"`
	// fee percentage
	FeePercentage string `json:"fee_percentage"`
	// tax flat
	TaxFlat string `json:"tax_flat"`
	// tax percentage
	TaxPercentage string `json:"tax_percentage"`
	// charge flat
	ChargeFlat string `json:"charge_flat"`
	// charge percentage
	ChargePercentage string `json:"charge_percentage"`
	// service flat
	ServiceFlat string `json:"service_flat"`
	// service percentage
	ServicePercentage string `json:"service_percentage"`
}
