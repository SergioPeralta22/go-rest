package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN string = "host=localhost port=5432 user=sergio dbname=realestate password=password sslmode=disable "
var DB *gorm.DB

func DBconnection() {
	var err error

	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB connection success")
	}
}
