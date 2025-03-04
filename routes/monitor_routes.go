package routes

import (
	"net/http"

	"brightpixel-backend/database"
	"brightpixel-backend/models"

	"github.com/gin-gonic/gin"
)

func GetMonitors(c *gin.Context) {
	var monitors []models.Monitor

	// Use sqlx Select function to retrieve all monitors
	err := database.DB.Select(&monitors, "SELECT * FROM monitors")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, monitors)
}

func SetupRoutes(r *gin.Engine) {
	r.GET("/monitors", GetMonitors)
}
