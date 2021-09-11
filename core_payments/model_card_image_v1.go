package core_payments

// Model to describe the card images that are retuned on the get call for retrieving users payment instrument.
type CardImageV1 struct {
	// The design associated with the card in the smallest dimension
	Small string `json:"small"`
	// The design associated with the card in the medium dimension
	Medium string `json:"medium"`
	// The design associated with the card in the large dimension
	Large string `json:"large"`
}
