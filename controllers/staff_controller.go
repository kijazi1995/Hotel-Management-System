package controllers

import (
	"Hotel-Management/config"
	"Hotel-Management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStaffs(c *gin.Context) {
	var staff []models.Staff
	if err := config.DB.Preload("Hotel").Find(&staff).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch staff"})
		return
	}
	c.JSON(http.StatusOK, staff)
}

func GetStaff(c *gin.Context) {
	id := c.Param("id")
	var staff models.Staff
	if err := config.DB.Preload("Hotel").First(&staff, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return
	}
	c.JSON(http.StatusOK, staff)
}

func CreateStaff(c *gin.Context) {
	var staff models.Staff
	if err := c.ShouldBindJSON(&staff); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := config.DB.Create(&staff).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create staff"})
		return
	}
	c.JSON(http.StatusCreated, staff)
}

func UpdateStaff(c *gin.Context) {
	id := c.Param("id")
	var staff models.Staff
	if err := config.DB.First(&staff, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return
	}

	if err := c.ShouldBindJSON(&staff); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := config.DB.Save(&staff).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update staff"})
		return
	}
	c.JSON(http.StatusOK, staff)
}

func DeleteStaff(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Staff{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete staff"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Staff deleted successfully"})
}
