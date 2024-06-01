package services

import (
	"context"
	"time"

	"github.com/IST0VE/site_pdf_api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *models.User) (*mongo.InsertOneResult, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return userCollection.InsertOne(ctx, user)
}

func UpdateUser(userID primitive.ObjectID, updateUser *models.User) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if updateUser.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateUser.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		updateUser.Password = string(hashedPassword)
	}

	update := bson.M{
		"$set": updateUser,
	}
	return userCollection.UpdateByID(ctx, userID, update)
}

func DeleteUser(userID primitive.ObjectID) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return userCollection.DeleteOne(ctx, bson.M{"_id": userID})
}

func GetUserByID(userID primitive.ObjectID) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	return &user, err
}

func GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	return &user, err
}
