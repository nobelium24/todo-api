package routes

import (
	"goWebService/pkg/controllers"
	"goWebService/pkg/services"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(router *gin.Engine) {
	todoService := services.NewTodoService()
	todoController := controllers.NewTodoController(todoService)

	// Define todo routes
	todoGroup := router.Group("/todos")
	{
		todoGroup.POST("/", todoController.CreateTodo)
		todoGroup.GET("/", todoController.GetAllTodos)
		todoGroup.GET("/:id", todoController.GetTodoById)
		todoGroup.PUT("/:id", todoController.UpdateTodo)
		todoGroup.DELETE("/:id", todoController.DeleteTodoById)
		todoGroup.DELETE("/user/:user_id", todoController.DeleteTodosByUserId) // New route for deleting todos by user ID
	}
}
