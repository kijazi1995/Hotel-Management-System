package controllers

import (
	"Hotel-Management/config"
	"Hotel-Management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGuests(c *gin.Context) {
	var guests []models.Guest
	if err := config.DB.Find(&guests).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to Fetch Gests"})
		return
	}
	c.JSON(http.StatusOK, guests)
}

func GetGuest(c *gin.Context) {
	id := c.Param("id")
	var guest models.Guest
	if err := config.DB.First(&guest, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guest not found"})
		return
	}
	c.JSON(http.StatusOK, guest)
}

func CreateGuest(c *gin.Context) {
	var guest models.Guest
	if err := c.ShouldBindJSON(&guest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := config.DB.Create(&guest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create guest"})
		return
	}
	c.JSON(http.StatusCreated, guest)
}

func UpdateGuest(c *gin.Context) {
	id := c.Param("id")
	var guest models.Guest
	if err := config.DB.First(&guest, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guest not found"})
		return
	}
	if err := c.ShouldBindJSON(&guest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := config.DB.Save(&guest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update guest"})
		return
	}
	c.JSON(http.StatusOK, guest)
}

func DeleteGuest(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Guest{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete guest"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Guest deleted successfully"})
}
