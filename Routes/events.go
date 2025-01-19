package routes

import (
	"net/http"
	"strconv"

	"example.com/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	var err error
    events, err := models.GetAllEvents()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Error to get events"})
        return
    }
    context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context){
	eventId,err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error to get events"})
        return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event id."})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context){
    var event models.Event
    err := context.ShouldBindJSON(&event)
    if err != nil{
        context.JSON(http.StatusBadRequest, gin.H{"message":"Cant parse Data!"})
        return
    }

	userId := context.GetInt64("userId")
    event.UserID = userId

    err = event.Save()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error to create event"})
		return
	}

    context.JSON(http.StatusCreated, gin.H{"message": "event created!", "event": event})
}

func updateEvent(context *gin.Context){ 
	var err error
	eventId,err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error to get events"})
        return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)



	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
        return
	} 

	if event.UserID != userId{
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error to get events"})
        return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error to Update event"})
        return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated"})
}

func deleteEvent(context *gin.Context){
	var err error
	eventId,err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error to get events"})
        return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
        return
	} 

	if event.UserID != userId{
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to delete this event"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event"})
        return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}
