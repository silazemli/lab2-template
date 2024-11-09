package reservation

import (
	"github.com/google/uuid"
)

type Hotel struct {
	ID       int       `json:"id"`
	HotelUID uuid.UUID `json:"uid"`
	Name     string    `json:"name"`
	Country  string    `json:"country"`
	City     string    `json:"city"`
	Address  string    `json:"address"`
	Stars    int       `json:"stars"`
	Price    int       `json:"price"`
}
