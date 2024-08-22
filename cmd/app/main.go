package main

import (
	"goWebService/pkg/routes"
	"goWebService/pkg/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	utils.ConnectToDB()
	routes.TodoRoutes(router)
	routes.UserRoutes(router)

	log.Println("Starting server on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
