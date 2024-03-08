package main

import (
	"conference-booking-rest-api/db"
	"conference-booking-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	err := server.Run(":8080") //localhost:8080
	if err != nil {
		return
	}
}
