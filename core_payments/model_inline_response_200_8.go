package core_payments
import (
	"time"
)

type InlineResponse2008 struct {
	// The value of the card security code.
	Value string `json:"value"`
	// The duration until which the code is valid given no transaction performed in that period
	Expiry time.Time `json:"expiry,omitempty"`
}
