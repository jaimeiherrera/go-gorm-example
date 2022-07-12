package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ConnectionString = "host=localhost port=5432 user=examples dbname=gormex password=Secret123"
var DB *gorm.DB

func DBConnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(ConnectionString), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to database")
	}

}
