package reservation

type hotelStorage interface {
	GetAll() ([]Hotel, error)
}

type reservationStorage interface {
	GetReservations(username string) ([]Reservation, error)
	GetReservation(reservationUID string) (Reservation, error)
	MakeReservation(reservation Reservation) error
	CancelReservation(reservationUID string) error
}
