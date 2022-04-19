package db

import (
	"log"

	"github.com/bernardn38/financefirst/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func ConnetDb() {
	dsn := "postgres://jxwsmvrzbtlahm:968955bbeb33eaf29375ad7f69a650544410349a772347c467734ee2c96072ba@ec2-34-194-158-176.compute-1.amazonaws.com:5432/danv5sg7puk1bd"
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
