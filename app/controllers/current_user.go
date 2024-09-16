package controllers

import (
	"net/http"
	"strings"

	"app/models"
	"app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CurrentUser(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing"})
		return
	}

	token = strings.TrimPrefix(token, "Bearer ")

	claims, err := utils.ParseJWT(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	var user models.User
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("email = ?", claims.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id": user.ID,
			"attributes": gin.H{
				"email":       user.Email,
				"name":        user.Name,
				"phone":       user.Phone,
				"postal_code": user.PostalCode,
				"street":      user.Street,
				"number":      user.Number,
				"complement":  user.Complement,
				"image_url":   "user.ImageURL",
			},
		},
	})
}
func UpdateUser(c *gin.Context) {
    userId := c.Param("id")

    token := c.GetHeader("Authorization")
    if token == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing"})
        return
    }

    token = strings.TrimPrefix(token, "Bearer ")

    claims, err := utils.ParseJWT(token)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
        return
    }

    if claims.Username != claims.Username {
        c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to update this user"})
        return
    }

    var userData models.User

    if err := c.Request.ParseMultipartForm(10 << 20); err != nil { 
        c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse form data"})
        return
    }

    userData.Email = c.Request.FormValue("user[email]")
    userData.Name = c.Request.FormValue("user[name]")
    userData.Phone = c.Request.FormValue("user[phone]")
    userData.PostalCode = c.Request.FormValue("user[postal_code]")
    userData.Street = c.Request.FormValue("user[street]")
    userData.Number = c.Request.FormValue("user[number]")
    userData.Complement = c.Request.FormValue("user[complement]")

    file, _, err := c.Request.FormFile("user[image]")
    if err == nil {
        _ = file 
    }

    db := c.MustGet("db").(*gorm.DB)
    var user models.User
    if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    if err := db.Model(&user).Updates(userData).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": user})
}