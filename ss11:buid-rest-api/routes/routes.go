package routes

import (
	midlewares "example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(server *gin.Engine) {
	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(midlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
