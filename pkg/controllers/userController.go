package controllers

import (
	"goWebService/pkg/models"
	"goWebService/pkg/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{userService: service}
}

func (tc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var existingUser models.User
	if err := tc.userService.GetUserByEmail(&existingUser, user.Email); err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	if err := tc.userService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "New user created"})
}

func (tc *UserController) GetAllUsers(c *gin.Context) {
	var users []models.User
	if err := tc.userService.GetAllUsers(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, users)
}

func (tc *UserController) GetSingleUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := tc.userService.GetUser(&user, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, user)
}

func (tc *UserController) GetUSerByEmail(c *gin.Context) {
	var user models.User
	var requestBody struct {
		Email string `json:"email"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	email := requestBody.Email
	if err := tc.userService.GetUserByEmail(&user, email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, user)
}

func (tc *UserController) UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	user.ID = uint(userId)

	if err := tc.userService.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (tc *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := tc.userService.DeleteUser(&user, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
