package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const DNS = "admin:password@tcp(127.0.0.1:3306)/db?parseTime=true"

func Connect() {
	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	db.AutoMigrate()

	DB = db
}
