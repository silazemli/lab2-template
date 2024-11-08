package server

import (
	"github.com/silazemli/lab2-template/internal/server/models"
)

type hotelStorage interface {
	GetAll() ([]hotel.Hotel, error)
}