package core_payments

type UsersUserIdBody struct {
	// Account state. Accepts values: `blocked`  `suspended`. Default is `suspended`.
	State string `json:"state,omitempty"`
}
