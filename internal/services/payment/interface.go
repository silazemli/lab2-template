package payment

type paymentStorage interface {
	PostPayment(price string) error
	CancelPayment(paymentUID string) error
}
