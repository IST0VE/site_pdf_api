package services

import (
	"log"

	"github.com/IST0VE/site_pdf_api/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func InitializeServices() {
	if config.DB == nil {
		log.Fatal("Database not initialized")
	}
	userCollection = config.DB.Collection("users")
}
