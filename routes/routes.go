package routes

import (
	"conference-booking-rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/conferences", getConferences)
	server.GET("/conferences/:id", getConference)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/conferences", createConference)
	authenticated.PUT("/conferences/:id", updateConference)
	authenticated.DELETE("/conferences/:id", deleteConference)
	authenticated.POST("/conferences/:id/register", registerForConference)
	authenticated.DELETE("/conferences/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
