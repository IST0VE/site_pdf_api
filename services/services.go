package services

import (
	"log"

	"github.com/IST0VE/site_pdf_api/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection
var apiTokenCollection *mongo.Collection
var packageCollection *mongo.Collection

func InitializeServices() {
	if config.DB == nil {
		log.Fatal("Database not initialized")
	}
	userCollection = config.DB.Collection("users")
	apiTokenCollection = config.DB.Collection("api_tokens")
	packageCollection = config.DB.Collection("packages")
}
