package reservation

import (
	"time"

	"github.com/google/uuid"
)

type Reservation struct {
	ID             int       `json:"id"`
	ReservationUID uuid.UUID `json:"reservationUid"`
	Username       string    `json:"username"`
	PaymentUID     uuid.UUID `json:"paymentUid"`
	HotelID        int       `json:"hotelId"`
	Status         string    `json:"status"`
	StartDate      time.Time `json:"startDate"`
	EndDate        time.Time `json:"endDate"`
}
