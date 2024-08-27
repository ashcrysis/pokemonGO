package middleware

import (
	"app/config"

	"github.com/gin-gonic/gin"
)

// InjectDB is a middleware that injects the database connection into the context
func InjectDB() gin.HandlerFunc {
    return func(c *gin.Context) {
        db := config.SetupDB()
        c.Set("db", db)
        c.Next()
    }
}
