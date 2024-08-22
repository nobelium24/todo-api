package controllers

import (
	"goWebService/pkg/models"
	"goWebService/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todoService *services.TodoService
}

func NewTodoController(service *services.TodoService) *TodoController {
	return &TodoController{todoService: service}
}

func (tc *TodoController) CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := tc.todoService.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

func (tc *TodoController) GetAllTodos(c *gin.Context) {
	var todos []models.Todo
	if err := tc.todoService.GetAllTodos(&todos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (tc *TodoController) GetTodoById(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	if err := tc.todoService.GetTodoByID(&todo, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (tc *TodoController) GetTodosByUserId(c *gin.Context) {
	var todo []models.Todo
	userId := c.Param("user_id")
	if err := tc.todoService.GetTodosByUserId(&todo, userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func (tc *TodoController) UpdateTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	if err := tc.todoService.GetTodoByID(&todo, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := tc.todoService.UpdateTodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, todo)
}

func (tc *TodoController) DeleteTodoById(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	if err := tc.todoService.DeleteTodoById(&todo, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}

func (tc *TodoController) DeleteTodosByUserId(c *gin.Context) {
	userID := c.Param("user_id")
	if err := tc.todoService.DeleteUserTodos(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todos deleted"})
}
