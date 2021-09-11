package core_payments
import (
	"time"
)

type UserFullV1Date struct {
	Registration time.Time `json:"registration,omitempty"`
}
