package core_payments

type CategoriesNameBody struct {
	// Serves as the maximum amount allowable for the category
	DefaultAmount string `json:"default_amount"`
	// Defines the default_amount type or context (UNIT, PERCENT). Defaults to UNIT
	Factor string `json:"factor,omitempty"`
	// If a comma-delimited list of Merchant Category Codes is passed, this category will use them as filters. Value with no record matches are ignored.
	Mcc string `json:"mcc,omitempty"`
	// A comma-delimited list of country abbreviations in [ISO 3166 alpha-3](http://www.iso.org/iso/country_codes) format. If provided, the currency of these countries will be used as filters. Value with no record matches are ignored.
	Country string `json:"country,omitempty"`
	// Spending priority ( The lower the number the hight the priority )
	Priority string `json:"priority,omitempty"`
	// If a comma-delimited list of card ID is passed, this category will use them as filters. Value with no record matches are ignored.
	Card string `json:"card,omitempty"`
	// If a comma-delimited list of Merchant Codes is passed, this category will use them as filters
	Merchant string `json:"merchant,omitempty"`
	// If a comma-delimited list of Terminal Codes is passed, this category will use them as filters
	Terminal string `json:"terminal,omitempty"`
}
