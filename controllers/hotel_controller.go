package controllers

import (
	"Hotel-Management/config"
	"Hotel-Management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHotels(c *gin.Context) {
	var hotels []models.Hotel
	if err := config.DB.Find(&hotels).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch hotels"})
		return
	}
	c.JSON(http.StatusOK, hotels)
}

func GetHotel(c *gin.Context) {
	id := c.Param("id")
	var hotel models.Hotel
	if err := config.DB.First(&hotel, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
		return
	}
	c.JSON(http.StatusOK, hotel)
}

func CreateHotel(c *gin.Context) {
	var hotel models.Hotel
	if err := c.ShouldBindJSON(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := config.DB.Create(&hotel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create hotel"})
		return
	}
	c.JSON(http.StatusCreated, hotel)
}

func UpdateHotel(c *gin.Context) {
	id := c.Param("id")
	var hotel models.Hotel
	if err := config.DB.First(&hotel, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
		return
	}
	if err := c.ShouldBindJSON(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	config.DB.Save(&hotel)
	c.JSON(http.StatusOK, hotel)
}

func DeleteHotel(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Hotel{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete hotel"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Hotel deleted successfully"})
}
