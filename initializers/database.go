package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var Db *gorm.DB

func ConnectDatabase() {
	var err error
	dbUrl := os.Getenv("DB_URL")
	Db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database")
	}
}
