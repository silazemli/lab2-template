package storage

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
