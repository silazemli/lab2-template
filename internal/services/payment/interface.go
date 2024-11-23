package payment

type paymentStorage interface {
	PostPayment(thePayment Payment) error
	CancelPayment(paymentUID string) error
}
