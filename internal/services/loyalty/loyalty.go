package models

type Loyalty struct {
	ID               int    `json:"id"`
	Username         string `json:"username"`
	ReservationCount int    `json:"reservation_count"`
	Status           string `json:"status"`
	Discount         int    `json:"discount"`
}
