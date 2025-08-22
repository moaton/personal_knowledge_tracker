package mongo

import (
	"context"
	"personal_knowledge_tracker/internal/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type fs struct {
	col *mongo.Collection
}

func NewFS(db *mongo.Database) *fs {
	return &fs{col: db.Collection("fs")}
}

func (f *fs) Save(ctx context.Context, file *entity.File, data []byte) error {
	return nil
}

func (f *fs) Get(ctx context.Context, id primitive.ObjectID) ([]byte, error) {
	return nil, nil
}

func (f *fs) Delete(ctx context.Context, id primitive.ObjectID) error {
	return nil
}
