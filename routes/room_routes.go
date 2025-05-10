package routes

import (
	"Hotel-Management/controllers"

	"github.com/gin-gonic/gin"
)

func RoomRoutes(r *gin.Engine) {
	roomGroup := r.Group("/rooms")
	{
		roomGroup.GET("/", controllers.GetRooms)
		roomGroup.GET("/:id", controllers.GetRoom)
		roomGroup.POST("/", controllers.CreateRoom)
		roomGroup.PUT("/:id", controllers.UpdateRoom)
		roomGroup.DELETE("/:id", controllers.DeleteRoom)
	}
}
