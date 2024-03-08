package routes

import (
	"conference-booking-rest-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getConferences(context *gin.Context) {
	conferences, err := models.GetAllConferences()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch conferences. Try again later."})
		return
	}
	context.JSON(http.StatusOK, conferences)
}

func getConference(context *gin.Context) {
	conferenceId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse conference id."})
		return
	}
	conference, err := models.GetConferenceByID(conferenceId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch conference."})
		return
	}
	context.JSON(http.StatusOK, conference)
}

func createConference(context *gin.Context) {
	var conference models.Conference
	err := context.ShouldBindJSON(&conference)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	conference.UserID = 1
	err = conference.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create conference. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Conference created!", "conference": conference})
}

func updateConference(context *gin.Context) {
	conferenceId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse conference id."})
		return
	}

	_, err = models.GetConferenceByID(conferenceId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch conference."})
		return
	}

	var updatedConference models.Conference

	err = context.ShouldBindJSON(&updatedConference)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	updatedConference.ID = conferenceId
	err = updatedConference.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update conference."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Conference updated successfully!", "conference": updatedConference})
}

func deleteConference(context *gin.Context) {
	conferenceId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse conference id."})
		return
	}

	conference, err := models.GetConferenceByID(conferenceId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the conference."})
		return
	}
	err = conference.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the conference."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Conference deleted successfully!"})
}
