package routes

import (
	"Hotel-Management/controllers"

	"github.com/gin-gonic/gin"
)

func HotelRoutes(r *gin.Engine) {
	hotelGroup := r.Group("/hotels")
	{
		hotelGroup.GET("/", controllers.GetHotels)
		hotelGroup.GET("/:id", controllers.GetHotel)
		hotelGroup.POST("/", controllers.GetHotel)
		hotelGroup.PUT("/:id", controllers.UpdateHotel)
		hotelGroup.DELETE("/:id", controllers.DeleteHotel)
	}
}
