package core_payments
import (
	"time"
)

type TransactionFundsTransferV1Date struct {
	Created time.Time `json:"created,omitempty"`
	Expiry time.Time `json:"expiry,omitempty"`
}
