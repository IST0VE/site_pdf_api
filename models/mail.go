package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TempMailbox struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Email     string             `bson:"email"`
	UserID    primitive.ObjectID `bson:"user_id"`
	CreatedAt time.Time          `bson:"created_at"`
	ExpiresAt time.Time          `bson:"expires_at"`
}

type Message struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	MailboxID  primitive.ObjectID `bson:"mailbox_id"`
	Content    string             `bson:"content"`
	ReceivedAt time.Time          `bson:"received_at"`
}
