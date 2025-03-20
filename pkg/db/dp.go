package db

import (
	"gorm.io/driver/postgres" 
	"gorm.io/gorm"
)

var DB *gorm.DB
func Connect() {
	dsn := "yourdatabase"
	db , err := gorm.Open(postgres.Open(dsn) , &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}

