package main

import (
	"github.com/gin-gonic/gin"
	"go-rest-api-crud/initializers"
	"go-rest-api-crud/routes"
	"os"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	r := gin.Default()

	routes.ToDoRoutes(r)

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
