package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/conferences", getConferences)
	server.GET("/conferences/:id", getConference)
	server.POST("/conferences", createConference)
	server.PUT("/conferences/:id", updateConference)
	server.DELETE("/conferences/:id", deleteConference)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
