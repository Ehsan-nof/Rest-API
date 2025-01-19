package routes

import (
	"net/http"
	"strconv"

	"example.com/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId,err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error to get events"})
        return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"messsage": "Could not fetch event"})
		return
	}

	err = event.Register(userId)

	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"messsage": "Could not register user for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"messsage": "Registered!"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId,err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error to get event"})
        return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"messsage": "Could not register user for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"messsage": "Cancelled!"})
}
