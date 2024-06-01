package services

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/IST0VE/site_pdf_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func generateRandomToken() string {
	token := make([]byte, 32) // 32 bytes = 256 bits
	if _, err := rand.Read(token); err != nil {
		return ""
	}
	return hex.EncodeToString(token)
}

func CreateAPIToken(apiToken *models.APIToken) (*mongo.InsertOneResult, error) {
	apiToken.Token = generateRandomToken() // Implement this function to generate a random token
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return apiTokenCollection.InsertOne(ctx, apiToken)
}

func GetAPITokenByID(tokenID primitive.ObjectID) (*models.APIToken, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var apiToken models.APIToken
	err := apiTokenCollection.FindOne(ctx, bson.M{"_id": tokenID}).Decode(&apiToken)
	return &apiToken, err
}

func GetAPITokenByToken(token string) (*models.APIToken, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var apiToken models.APIToken
	err := apiTokenCollection.FindOne(ctx, bson.M{"token": token}).Decode(&apiToken)
	return &apiToken, err
}

func GetAPITokensByUserID(userID primitive.ObjectID) ([]*models.APIToken, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var apiTokens []*models.APIToken
	cursor, err := apiTokenCollection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var apiToken models.APIToken
		if err := cursor.Decode(&apiToken); err != nil {
			return nil, err
		}
		apiTokens = append(apiTokens, &apiToken)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return apiTokens, nil
}

func UpdateAPIToken(tokenID primitive.ObjectID, updateToken *models.APIToken) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{}

	if updateToken.Name != "" {
		update["name"] = updateToken.Name
	}
	if updateToken.TotalRequests != 0 {
		update["total_requests"] = updateToken.TotalRequests
	}
	if updateToken.RemainingRequests != 0 {
		update["remaining_requests"] = updateToken.RemainingRequests
	}
	if updateToken.UserID != primitive.NilObjectID {
		update["user_id"] = updateToken.UserID
	}

	return apiTokenCollection.UpdateByID(ctx, tokenID, bson.M{"$set": update})
}

func DeleteAPIToken(tokenID primitive.ObjectID) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return apiTokenCollection.DeleteOne(ctx, bson.M{"_id": tokenID})
}
