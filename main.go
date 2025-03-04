package main

import (
	"brightpixel-backend/database"
	"brightpixel-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB() // Initialize database connection

	r := gin.Default()
	routes.SetupRoutes(r) // Register routes

	r.Run(":8080") // Start the server
}
