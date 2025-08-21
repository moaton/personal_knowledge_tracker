package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Resource struct {
	ID        primitive.ObjectID     `bson:"_id,omitempty"`
	Title     string                 `bson:"title"`
	Type      string                 `bson:"type"`
	Content   string                 `bson:"content"`
	Tags      []string               `bson:"tags"`
	Metadata  map[string]interface{} `bson:"metadata"`
	CreatedAt time.Time              `bson:"created_at"`
	UpdatedAt time.Time              `bson:"updated_at"`
}
