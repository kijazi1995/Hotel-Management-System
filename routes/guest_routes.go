package routes

import (
	"Hotel-Management/controllers"

	"github.com/gin-gonic/gin"
)

func GuestRoutes(r *gin.Engine) {
	guestGroup := r.Group("/guests")
	{
		guestGroup.GET("/", controllers.GetGuests)
		guestGroup.GET("/:id", controllers.GetGuest)
		guestGroup.POST("/", controllers.CreateGuest)
		guestGroup.PUT("/:id", controllers.UpdateGuest)
		guestGroup.DELETE("/:id", controllers.DeleteGuest)
	}
}
