// routes/user_routes.go
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"brightpixel-backend/database"
	"brightpixel-backend/models"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User
	err := database.DB.Select(&users, "SELECT * FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
