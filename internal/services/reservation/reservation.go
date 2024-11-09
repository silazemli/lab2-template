package reservation

import (
	"time"

	"github.com/google/uuid"
)

type Reservation struct {
	ID             int       `json:"id"`
	ReservationUID uuid.UUID `json:"reservation_uid"`
	Username       string    `json:"username"`
	PaymentUID     uuid.UUID `json:"payment_uid"`
	HotelID        int       `json:"hotel_id"`
	Status         string    `json:"status"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
}
