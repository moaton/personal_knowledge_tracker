package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	ResourceID primitive.ObjectID `bson:"resource_id"`
	UserID     primitive.ObjectID `bson:"user_id"`
	Rating     int                `bson:"rating"`
	Comment    string             `bson:"comment"`
}
