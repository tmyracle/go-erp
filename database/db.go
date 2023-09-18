package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	Err error
)


func InitDatabase() {
	dsn := os.Getenv("DATABASE_DSN")
	DB, Err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if Err != nil {
		panic("failed to connect database")
	}
}