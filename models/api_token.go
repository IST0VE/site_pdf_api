package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type APIToken struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Token             string             `bson:"token" json:"token"`
	Name              string             `bson:"name" json:"name"`
	UserID            primitive.ObjectID `bson:"user_id" json:"user_id"`
	TotalRequests     int                `bson:"total_requests" json:"total_requests"`
	RemainingRequests int                `bson:"remaining_requests" json:"remaining_requests"`
}
