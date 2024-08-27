package main

import (
	"app/config"
	"app/controllers"
	"app/middleware"
	"app/models"

	"github.com/gin-gonic/gin"
)

func main() {
    // Initialize the database connection
    db := config.SetupDB()

    // Auto migrate the user model
    db.AutoMigrate(&models.User{})

    r := gin.Default()

    // Apply the InjectDB middleware
    r.Use(middleware.InjectDB())

    // Define the routes
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)
	r.POST("/logoff", controllers.Logoff) 
    r.GET("/protected", middleware.AuthenticateJWT(), controllers.Protected)

    // Start the server
    r.Run(":8080")
}