package payment

import (
	"github.com/google/uuid"
)

type Payment struct {
	ID         int       `json:"id"`
	PaymentUID uuid.UUID `json:"uid"`
	Status     string    `json:"status"`
	Price      int       `json:"price"`
}
