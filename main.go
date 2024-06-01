package main

import (
	"github.com/IST0VE/site_pdf_api/config"
	"github.com/IST0VE/site_pdf_api/handlers"
	"github.com/IST0VE/site_pdf_api/middlewares"
	"github.com/IST0VE/site_pdf_api/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	config.SetupDatabase()

	// Initialize RabbitMQ connection
	config.SetupRabbitMQ()

	// Initialize services
	services.InitializeServices()

	r := gin.Default()

	// Public routes
	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.LoginUser)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.PUT("/users/:id", handlers.UpdateUser)
		protected.DELETE("/users/:id", handlers.DeleteUser)

		// API Token routes
		protected.POST("/api-tokens", handlers.CreateAPIToken)
		protected.GET("/api-tokens/:id", handlers.GetAPIToken)
		protected.GET("/api-tokens/user/:user_id", handlers.GetAPITokensByUserID)
		protected.PUT("/api-tokens/:id", handlers.UpdateAPIToken)
		protected.DELETE("/api-tokens/:id", handlers.DeleteAPIToken)

		// Package routes
		protected.POST("/packages", handlers.CreatePackage)
		protected.GET("/packages", handlers.GetAllPackages)
		protected.GET("/packages/:id", handlers.GetPackage)
		protected.PUT("/packages/:id", handlers.UpdatePackage)
		protected.DELETE("/packages/:id", handlers.DeletePackage)
	}

	r.Run(":8080")
}
