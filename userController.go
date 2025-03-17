package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rajritwika1/codwikz/database"
	"github.com/rajritwika1/codwikz/models"
)

// Get User Profile (Example Function)
func GetUserProfile(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	result := database.DB.First(&user, userID)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Update User Profile (Example Function)
func UpdateUserProfile(c *gin.Context) {
	userID := c.Param("id")
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Model(&models.User{}).Where("id = ?", userID).Updates(user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully!"})
}
