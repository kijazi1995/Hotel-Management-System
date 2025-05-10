package controllers

import (
	"Hotel-Management/config"
	"Hotel-Management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRoomServices(c *gin.Context) {
	var services []models.RoomService
	if err := config.DB.Preload("Guest").Preload("Room").Find(&services).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch room services"})
		return
	}
	c.JSON(http.StatusOK, services)
}

func GetRoomService(c *gin.Context) {
	id := c.Param("id")
	var service models.RoomService
	if err := config.DB.Preload("Guest").Preload("Room").First(&service, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room service not found"})
		return
	}
	c.JSON(http.StatusOK, service)
}

func CreateRoomService(c *gin.Context) {
	var service models.RoomService
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := config.DB.Create(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create room service"})
		return
	}
	c.JSON(http.StatusCreated, service)
}

func UpdateRoomService(c *gin.Context) {
	id := c.Param("id")
	var service models.RoomService
	if err := config.DB.First(&service, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room service not found"})
		return
	}
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := config.DB.Save(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update room service"})
		return
	}
	c.JSON(http.StatusOK, service)
}

func DeleteRoomService(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.RoomService{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete room service"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Room service deleted successfully"})
}
