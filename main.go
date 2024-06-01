package main

import (
	"github.com/IST0VE/site_pdf_api/config"
	"github.com/IST0VE/site_pdf_api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize database connection
	config.SetupDatabase()

	// Initialize RabbitMQ connection
	config.SetupRabbitMQ()

	// User routes
	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.LoginUser)
	r.PUT("/users/:id", handlers.UpdateUser)

	r.Run(":8080")
}
