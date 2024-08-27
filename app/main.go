package main

import (
	"app/config"
	"app/controllers"
	"app/middleware"
	"app/models"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := config.SetupDB()

	db.AutoMigrate(&models.User{})

	r := gin.Default()

	r.Use(middleware.InjectDB())

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/logoff", controllers.Logoff)
	r.GET("/protected", middleware.AuthenticateJWT(), controllers.Protected)

	r.Run(":8080")
}
