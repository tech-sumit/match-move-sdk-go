package core_payments
import (
	"time"
)

type UserIdentificationV1 struct {
	Type_ string `json:"type,omitempty"`
	Number string `json:"number,omitempty"`
	Issue time.Time `json:"issue,omitempty"`
	Expiry time.Time `json:"expiry,omitempty"`
}
