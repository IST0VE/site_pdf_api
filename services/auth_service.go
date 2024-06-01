package services

import (
	"context"
	"time"

	"./config"
	"./models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userConnection *mongo.Collection = config.DB.Collection("users")

// RegisterUser создает нового пользователя
func RegisterUser(user *models.User) (*mongo.InsertOneResult, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return userCollection.InsertOne(ctx, user)
}

// AuthenticateUser аутентифицирует пользователя и генерирует JWT
func AuthenticateUser(email, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	return GenerateJWT(email)
}
