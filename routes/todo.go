package routes

import (
	"github.com/gin-gonic/gin"
	"go-rest-api-crud/controllers"
)

func ToDoRoutes(r *gin.Engine) {
	userGroup := r.Group("/toDo")

	{
		userGroup.POST("", controllers.ToDoCreate)
		userGroup.GET("", controllers.ToDoIndex)
		userGroup.GET("/:id", controllers.ToDoGet)
		userGroup.PUT("/:id", controllers.ToDoUpdate)
		userGroup.DELETE("/:id", controllers.ToDoDelete)
	}
}
