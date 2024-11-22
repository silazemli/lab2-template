package payment

type Payment struct {
	ID         int    `json:"id"`
	PaymentUID string `json:"uid"`
	Status     string `json:"status"`
	Price      int    `json:"price"`
}
