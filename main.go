package main

import (
	"github.com/IST0VE/site_pdf_api/config"
	"github.com/IST0VE/site_pdf_api/handlers"
	"github.com/IST0VE/site_pdf_api/middlewares"
	"github.com/IST0VE/site_pdf_api/services"
	"github.com/gin-gonic/gin"
)

func main() {

	config.SetupDatabase()

	config.SetupRabbitMQ()

	services.InitializeServices()

	r := gin.Default()

	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.LoginUser)

	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.PUT("/users/:id", handlers.UpdateUser)
	}

	r.Run(":8080")
}
