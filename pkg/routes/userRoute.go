package routes

import (
	"goWebService/pkg/controllers"
	"goWebService/pkg/middlewares"
	"goWebService/pkg/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)

	userGroup := router.Group("/users")
	{
		userGroup.POST("/", userController.CreateUser)
		userGroup.GET("/", middlewares.AuthMiddleware(), userController.GetAllUsers)
		userGroup.GET("/user", middlewares.AuthMiddleware(), userController.GetSingleUser)
		userGroup.POST("/email", middlewares.AuthMiddleware(), userController.GetUSerByEmail)
		userGroup.PUT("/", middlewares.AuthMiddleware(), userController.UpdateUser)
		userGroup.DELETE("/", middlewares.AuthMiddleware(), userController.DeleteUser)
	}
}
