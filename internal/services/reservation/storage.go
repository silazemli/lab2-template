package reservation

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type storage struct {
	db *gorm.DB
}

func NewDB() (*storage, error) {
	dsn := os.Getenv("PGDSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return &storage{}, err
	}
	return &storage{db}, nil
}

func (stg *storage) GetAll() ([]Hotel, error) {
	hotels := []Hotel{}
	err := stg.db.Table("hotels").Find(&hotels).Error
	if err != nil {
		return []Hotel{}, err
	}
	return hotels, nil
}

func (stg *storage) GetReservations(username string) ([]Reservation, error) {
	reservations := []Reservation{}
	err := stg.db.Table("users").Where("username = ?", username).Find(&reservations).Error
	if err != nil {
		return []Reservation{}, err
	}
	return reservations, nil
}

func (stg *storage) GetReservation(reservationUID string) (Reservation, error) {
	reservation := Reservation{}
	err := stg.db.Table("reservations").Where("uuid = ?", reservationUID).Take(&reservation).Error
	if err != nil {
		return Reservation{}, err
	}
	return reservation, err
}

func (stg *storage) MakeReservation(reservation Reservation) error {
	err := stg.db.Table("reservations").Create(&reservation).Error
	if err != nil {
		return err
	}
	return nil
}

func (stg *storage) CancelReservation(reservationUID string) error {
	reservation := Reservation{}
	err := stg.db.Table("reservations").Where("uid = ?", reservationUID).Take(&reservation).Error
	if err != nil {
		return err
	}
	reservation.Status = "CANCELED"
	err = stg.db.Table("reservations").Where("uid = ?", reservationUID).Updates(&reservation).Error
	if err != nil {
		return err
	}
	return nil
}
