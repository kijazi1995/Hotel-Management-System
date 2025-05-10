package routes

import (
	"Hotel-Management/controllers"

	"github.com/gin-gonic/gin"
)

func BookingRoutes(r *gin.Engine) {
	guestGroup := r.Group("/guests")
	{
		guestGroup.GET("/", controllers.GetBookings)
		guestGroup.GET("/:id", controllers.GetBooking)
		guestGroup.POST("/", controllers.CreateBooking)
		guestGroup.PUT("/:id", controllers.UpdateBooking)
		guestGroup.DELETE("/:id", controllers.DeleteBooking)
	}
}
