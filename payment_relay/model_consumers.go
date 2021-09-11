package payment_relay

type Consumers struct {
	// Consumer Hash ID
	ConsumerHash string `json:"consumer_hash"`
	// Consumer ACK Key
	AclKey string `json:"acl_key"`
	// Consumer code name
	Code string `json:"code"`
	// Consumer description
	Description string `json:"description"`
	// Will call a client's external pre-credit url before doing payment and topup [`0`, `1`]
	PreCreditConfirm string `json:"pre_credit_confirm,omitempty"`
	// Client's external pre-credit url method to call before doing payment and topup [`GET`, `POST`, `PUT`, `DELETE`]
	PreCreditUrlMethod string `json:"pre_credit_url_method,omitempty"`
	// Client's external pre-credit url to call before doing payment and topup
	PreCreditUrl string `json:"pre_credit_url,omitempty"`
	// Will call a client's external pre-credit url after doing payment and topup [`0`, `1`]
	PostCreditConfirm string `json:"post_credit_confirm,omitempty"`
	// Client's external pre-credit url method to call after doing payment and topup [`GET`, `POST`, `PUT`, `DELETE`]
	PostCreditUrlMethod string `json:"post_credit_url_method,omitempty"`
	// Client's external pre-credit url to call after doing payment and topup
	PostCreditUrl string `json:"post_credit_url,omitempty"`
	// Timezone in [TZ format](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)
	Timezone string `json:"timezone,omitempty"`
	// Time Offset in [UTC format](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)
	TimeOffset string `json:"time_offset,omitempty"`
	// Status [`active`, `inactive`, `terminated`]
	Status string `json:"status"`
	// Is admin
	IsAdmin int32 `json:"is_admin"`
	// Date created `(YYYY-MM-DD HH:MM:SS)`
	DateCreated string `json:"date_created,omitempty"`
	// Date modified `(YYYY-MM-DD HH:MM:SS)`
	DateModified string `json:"date_modified,omitempty"`
}
