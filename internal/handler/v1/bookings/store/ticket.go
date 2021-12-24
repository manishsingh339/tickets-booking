package store

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ticket struct {
	ID          primitive.ObjectID `bson:"_id"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	EventTime   time.Time          `bson:"event_time"`
	EventTitle  string             `bson:"event_title"`
	Description string             `bson:"description"`
}
