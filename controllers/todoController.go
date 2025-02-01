package controllers

import (
	"github.com/gin-gonic/gin"
	"go-rest-api-crud/initializers"
	"go-rest-api-crud/models"
)

func ToDoCreate(c *gin.Context) {
	var body struct {
		Content string
		Status  bool
	}
	c.Bind(&body)
	todo := models.ToDo{Content: body.Content, Status: body.Status}
	result := initializers.Db.Create(todo)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"toDo": todo,
	})
}

func ToDoIndex(c *gin.Context) {
	var toDos []models.ToDo
	initializers.Db.Find(&toDos)

	c.JSON(200, gin.H{
		"toDos": toDos,
	})
}

func ToDoShow(c *gin.Context) {
	id := c.Param("id")

	var todo models.ToDo
	initializers.Db.First(&todo, id)

	c.JSON(200, gin.H{
		"toDo": todo,
	})
}

func ToDoUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Content string
		Status  bool
	}
	c.Bind(&body)

	var todo models.ToDo
	initializers.Db.First(&todo, id)

	initializers.Db.Model(&todo).Updates(models.ToDo{
		Content: body.Content,
		Status:  body.Status,
	})

	c.JSON(200, gin.H{
		"toDo": todo,
	})
}

func ToDoDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.Db.Delete(&models.ToDo{}, id)

	c.JSON(200, gin.H{
		"message": "Todo" + id + "removed successfully!",
	})

}
