package main

import (
	"app/config"
	"app/controllers"
	"app/middleware"
	"app/models"
	"log"
	"time"

	"github.com/gin-contrib/cors"

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

	config := cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
        ExposeHeaders:    []string{"Authorization"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }
	
    r.Use(cors.New(config))
	
	r.Use(middleware.InjectDB())
	pokemonController := controllers.NewPokemonController()
	pokemonRoutes := r.Group("v2/pokemons")
	{
		pokemonRoutes.GET("/search", pokemonController.Search)      
		pokemonRoutes.GET("/fetch_all", pokemonController.FetchAllPokemonData) 
		pokemonRoutes.GET("/species", pokemonController.Species)     
		pokemonRoutes.GET("/toggle_api", pokemonController.ToggleAPI)   
	}
	v2 := r.Group("/v2")
    {
		users := v2.Group("/users")
        {
			users.GET("/current", middleware.AuthenticateJWT(), controllers.CurrentUser)
			users.PUT("/update/:id", middleware.AuthenticateJWT(), controllers.UpdateUser)
		}
    }
	r.POST("/signup", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/logoff", controllers.Logoff)
	r.GET("/protected", middleware.AuthenticateJWT(), controllers.Protected)
	
	r.Run(":8080")
}
