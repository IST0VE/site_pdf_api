package config

import (
	"context"
	"log"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var RabbitConn *amqp.Connection

func SetupDatabase() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://nicorovko:Bj3whCrfp1aVZYK4@pdfsite.pfpvece.mongodb.net/?retryWrites=true&w=majority&appName=pdfSite")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database("pdfSite")
}

func SetupRabbitMQ() {
	var err error
	RabbitConn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
}
