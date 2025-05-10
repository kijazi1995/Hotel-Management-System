package controllers

import (
	"Hotel-Management/config"
	"Hotel-Management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAmenities(c *gin.Context) {
	var amenities []models.Amenity
	if err := config.DB.Find(&amenities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch amenities"})
		return
	}
	c.JSON(http.StatusOK, amenities)
}

func GetAmenity(c *gin.Context) {
	id := c.Param("id")
	var amenity models.Amenity
	if err := config.DB.First(&amenity, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Amenity not found"})
		return
	}
	c.JSON(http.StatusOK, amenity)
}

func CreateAmenity(c *gin.Context) {
	var amenity models.Amenity
	if err := c.ShouldBindJSON(&amenity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := config.DB.Create(&amenity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create amenity"})
		return
	}
	c.JSON(http.StatusCreated, amenity)
}
func UpdateAmenity(c *gin.Context) {
	id := c.Param("id")
	var amenity models.Amenity
	if err := config.DB.First(&amenity, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Amenity not found"})
		return
	}
	if err := c.ShouldBindJSON(&amenity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := config.DB.Save(&amenity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update amenity"})
		return
	}
	c.JSON(http.StatusOK, amenity)
}

func DeleteAmenity(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Amenity{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete amenity"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Amenity deleted successfully"})
}
