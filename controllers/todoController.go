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
	result := initializers.Db.Create(&todo)

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

func ToDoGet(c *gin.Context) {
	id := c.Param("id")

	var todo models.ToDo
	initializers.Db.First(&todo, id)

	if todo.ID == 0 {
		c.JSON(400, gin.H{
			"message": "Entity " + id + " not found!",
		})
		return
	}

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

	var todo models.ToDo
	initializers.Db.First(&todo, id)

	if todo.Content == "" {
		c.JSON(500, gin.H{
			"message": "Already deleted!",
		})
		return
	}

	initializers.Db.Delete(&models.ToDo{}, id)

	c.JSON(200, gin.H{
		"message": "Todo " + id + " removed successfully!",
	})

}
