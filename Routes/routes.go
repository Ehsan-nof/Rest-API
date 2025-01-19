package routes

import (
	"example.com/middlewares"
	"github.com/gin-gonic/gin"
)
func Register(server *gin.Engine){
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	
    server.POST("/events", middlewares.Authenticate, createEvent)
	server.PUT("/events/:id", middlewares.Authenticate,updateEvent)
	server.DELETE("/events/:id", middlewares.Authenticate,deleteEvent)

	server.POST("/events/:id/register", middlewares.Authenticate, registerForEvent)
	server.DELETE("/events/:id/register", middlewares.Authenticate, cancelRegistration)

	server.POST("/signup",signup)
	server.POST("/login",login)
}
