package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "host= user= password= dbname= sslmode=disable"

func ConnectDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database")
	}

	fmt.Println("Database Connected")
	return db
}
