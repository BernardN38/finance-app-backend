package db

import (
	"log"

	"github.com/bernardn38/financefirst/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func ConnetDb() {
	dsn := "host=localhost user=eris dbname=finance_app port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("could not open connection to database")
	}
	DBConn = db
}

func InitDb() {
	ConnetDb()
	DBConn.AutoMigrate(&models.Transactions{}, &models.User{})
	log.Println("Database Migrated")
}
