package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	To        string             `bson:"to"`
	Content   string             `bson:"content"`
	Sent      bool               `bson:"sent"`
	SentAt    time.Time          `bson:"sent_at,omitempty"`
	MessageID string             `bson:"message_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
}
