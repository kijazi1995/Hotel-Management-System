package main

import (
	"Hotel-Management/config"
	"Hotel-Management/routes"

	//"Hotel-Management/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize DB
	config.ConnectDB()

	// Initialize router
	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r)

	// Load HTML templates
	r.LoadHTMLGlob("Hotel-Management-ui/*.html")

	//routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// Run server
	r.Run(":8080")
}
