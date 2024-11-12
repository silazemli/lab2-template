package loyalty

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

func (stg *storage) GetStatus(username string) (string, error) {
	loyalty := Loyalty{}
	err := stg.db.Table("users").Where("username = ?", username).Take(&loyalty).Error
	if err != nil {
		return loyalty.Status, err
	}
	return loyalty.Status, nil
}

func (stg *storage) IncrementCounter(username string) error {
	loyalty := Loyalty{}
	err := stg.db.Table("users").Where("username = ?", username).Take(&loyalty).Error
	if err != nil {
		return err
	}
	loyalty.ReservationCount += 1
	loyalty.Status = GetStatus(loyalty.ReservationCount)
	err = stg.db.Table("users").Where("username = ?", username).Updates(&loyalty).Error
	if err != nil {
		return err
	}
	return nil
}

func (stg *storage) DecrementCounter(username string) error {
	loyalty := Loyalty{}
	err := stg.db.Table("users").Where("username = ?", username).Take(&loyalty).Error
	if err != nil {
		return err
	}
	loyalty.ReservationCount -= 1
	loyalty.Status = GetStatus(loyalty.ReservationCount)
	err = stg.db.Table("users").Where("username = ?", username).Updates(&loyalty).Error
	if err != nil {
		return err
	}
	return nil
}

func GetStatus(numberOfReservations int) string {
	if numberOfReservations >= 20 {
		return "GOLD"
	} else if numberOfReservations >= 10 {
		return "SILVER"
	} else {
		return "BRONZE"
	}
}
