package payment

import (
	"github.com/google/uuid"
)

type paymentStorage interface {
	PostPayment(price int) error
	CancelPayment(paymentUID uuid.UUID) error
}
