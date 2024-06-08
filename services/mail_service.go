package services

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/IST0VE/site_pdf_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var mailboxCollection *mongo.Collection // Инициализация при старте приложения

func GenerateUniqueEmail() (string, error) {
	rand.Seed(time.Now().UnixNano())
	uniquePart := rand.Int63()
	timestamp := time.Now().Unix()
	email := fmt.Sprintf("%d%d@example.com", uniquePart, timestamp)
	return email, nil
}

func CreateMailbox(email string, userID primitive.ObjectID) models.TempMailbox {
	now := time.Now()
	expires := now.Add(10 * time.Minute) // Действует 10 минут
	mailbox := models.TempMailbox{
		Email:     email,
		UserID:    userID,
		CreatedAt: now,
		ExpiresAt: expires,
	}
	// Сохранение в базу данных
	mailboxCollection.InsertOne(context.Background(), mailbox)
	return mailbox
}

var messageCollection *mongo.Collection // Это должно быть инициализировано где-то в вашем коде

func GetMessages(mailboxID string) ([]models.Message, error) {
	var messages []models.Message
	filter := bson.M{"mailbox_id": mailboxID}
	cursor, err := messageCollection.Find(context.Background(), filter)
	if err != nil {
		log.Println("Error finding messages:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var message models.Message
		if err := cursor.Decode(&message); err != nil {
			log.Println("Error decoding message:", err)
			return nil, err
		}
		messages = append(messages, message)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Error with cursor:", err)
		return nil, err
	}

	return messages, nil
}
