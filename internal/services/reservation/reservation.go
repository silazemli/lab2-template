package reservation

import (
	"time"
)

type Reservation struct {
	ReservationUID string    `json:"reservation_uid"`
	Username       string    `json:"username"`
	PaymentUID     string    `json:"payment_uid"`
	HotelID        int       `json:"hotel_id"`
	Status         string    `json:"status"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
}
