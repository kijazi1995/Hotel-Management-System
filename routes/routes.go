package routes

import (
	"Hotel-Management/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	//Hotel routes
	r.GET("/hotels", controllers.GetHotels)
	r.GET("hotel/:id", controllers.GetHotel)
	r.POST("/hotel", controllers.CreateHotel)
	r.PUT("hotel/:id", controllers.UpdateHotel)
	r.DELETE("hotel/:id", controllers.DeleteHotel)

	//Room routes
	r.GET("/rooms", controllers.GetRooms)
	r.GET("room/:id", controllers.GetRoom)
	r.POST("/room", controllers.CreateRoom)
	r.PUT("room/:id", controllers.UpdateRoom)
	r.DELETE("room/:id", controllers.DeleteRoom)

	//guest routes
	r.GET("/guests", controllers.GetGuests)
	r.GET("guest/:id", controllers.GetGuest)
	r.POST("/guest", controllers.CreateGuest)
	r.PUT("guest/:id", controllers.UpdateGuest)
	r.DELETE("guest/:id", controllers.DeleteGuest)

	//Bookings routes
	r.GET("/bookings", controllers.GetBookings)
	r.GET("booking/:id", controllers.GetBooking)
	r.POST("/booking", controllers.CreateBooking)
	r.PUT("booking/:id", controllers.UpdateBooking)
	r.DELETE("booking/:id", controllers.DeleteBooking)

	//Payments routes
	r.GET("/payments", controllers.GetPayments)
	r.GET("payment/:id", controllers.GetPayment)
	r.POST("/payment", controllers.CreatePayment)
	r.PUT("payment/:id", controllers.UpdatePayment)
	r.DELETE("payment/:id", controllers.DeletePayment)

	//Amenities routes
	r.GET("/amenities", controllers.GetAmenities)
	r.GET("amenity/:id", controllers.GetAmenity)
	r.POST("/amenity", controllers.CreateAmenity)
	r.PUT("amenity/:id", controllers.UpdateAmenity)
	r.DELETE("amenity/:id", controllers.DeleteAmenity)

	//Staffs  routes
	r.GET("/staffs", controllers.GetStaffs)
	r.GET("staff/:id", controllers.GetStaff)
	r.POST("/staff", controllers.CreateStaff)
	r.PUT("staff/:id", controllers.UpdateStaff)
	r.DELETE("staff/:id", controllers.DeleteStaff)

	//review  routes
	r.GET("/reviews", controllers.GetReviews)
	r.GET("review/:id", controllers.GetReview)
	r.POST("/review", controllers.CreateReview)
	r.PUT("review/:id", controllers.UpdateReview)
	r.DELETE("review/:id", controllers.DeleteReview)

	//room services  routes
	r.GET("/room_services", controllers.GetRoomServices)
	r.GET("room_service/:id", controllers.GetRoomService)
	r.POST("/room_service", controllers.CreateRoomService)
	r.PUT("room_service/:id", controllers.UpdateRoomService)
	r.DELETE("reroom_serviceview/:id", controllers.DeleteRoomService)
	// HotelRoutes(r)
	// RoomRoutes(r)
	// GuestRoutes(r)
	// BookingRoutes(r)
	// Add more routes here later (Rooms, Guests, etc.)
}
