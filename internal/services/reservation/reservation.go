package reservation

import (
	"time"
)

type Reservation struct {
	ID             int       `json:"id"`
	ReservationUID string    `json:"reservationUid"`
	Username       string    `json:"username"`
	PaymentUID     string    `json:"paymentUid"`
	HotelUID       string    `json:"hotelUid"`
	Status         string    `json:"status"`
	StartDate      time.Time `json:"startDate"`
	EndDate        time.Time `json:"endDate"`
}
