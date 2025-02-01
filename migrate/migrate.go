package main

import (
	"go-rest-api-crud/initializers"
	"go-rest-api-crud/models"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	err := initializers.Db.AutoMigrate(&models.ToDo{})
	if err != nil {
		log.Fatal("Error migrating database")
	}
}
