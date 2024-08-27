package middleware

import (
	"app/config"

	"github.com/gin-gonic/gin"
)

func InjectDB() gin.HandlerFunc {
    return func(c *gin.Context) {
        db := config.SetupDB()
        c.Set("db", db)
        c.Next()
    }
}
