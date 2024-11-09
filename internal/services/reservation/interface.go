package reservation

import (
	"github.com/google/uuid"
)

type reservationStorage interface {
	GetAll() ([]Hotel, error)
	GetReservations(username string) ([]Reservation, error)
	GetReservation(reservationUID uuid.UUID)
}
