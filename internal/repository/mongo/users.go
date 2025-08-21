package mongo

import (
	"context"
	"personal_knowledge_tracker/internal/entity"

	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	col *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *userRepository {
	return &userRepository{col: db.Collection("users")}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	_, err := r.col.InsertOne(ctx, user)
	return err
}
