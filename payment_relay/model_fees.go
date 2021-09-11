package payment_relay

type Fees struct {
	// Fees Hash Id
	Id string `json:"id"`
	Consumer *FeesConsumer `json:"consumer"`
	Product *FeesProduct `json:"product"`
	Provider *FeesProvider `json:"provider"`
	ProviderChannel *FeesProviderChannel `json:"provider_channel"`
	// Basic fee in flat
	FeeFlat float64 `json:"fee_flat"`
	// Basic fee in percentage
	FeePercentage float64 `json:"fee_percentage"`
	// Tax fee in flat
	TaxFlat float64 `json:"tax_flat"`
	// Tax fee in percentage
	TaxPercentage float64 `json:"tax_percentage"`
	// Charge fee in flat
	ChargeFlat float64 `json:"charge_flat"`
	// Charge fee in percentage
	ChargePercentage float64 `json:"charge_percentage"`
	// Service fee in flat
	ServiceFlat float64 `json:"service_flat"`
	// Service fee in percentage
	ServicePercentage float64 `json:"service_percentage"`
	// Is active
	IsActive int32 `json:"is_active"`
	// Date inactivated
	DateInactive string `json:"date_inactive,omitempty"`
	ModifiedBy *FeesModifiedBy `json:"modified_by,omitempty"`
	// Date created
	DateCreated string `json:"date_created"`
	// Date modified
	DateModified string `json:"date_modified"`
}
