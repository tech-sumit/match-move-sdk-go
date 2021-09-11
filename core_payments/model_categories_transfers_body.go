package core_payments

type CategoriesTransfersBody struct {
	// Amount to be transferred
	Amount string `json:"amount"`
	// Fund category name where the amount will be originating.
	From string `json:"from"`
	// Fund category name where the amount will be transferred.
	To string `json:"to"`
}
