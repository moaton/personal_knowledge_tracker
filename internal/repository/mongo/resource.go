package mongo

import (
	"context"
	"personal_knowledge_tracker/internal/entity"

	"go.mongodb.org/mongo-driver/mongo"
)

type resourceRepository struct {
	col *mongo.Collection
}

func NewResourceRepository(db *mongo.Database) *resourceRepository {
	return &resourceRepository{col: db.Collection("resources")}
}

func (r *resourceRepository) Create(ctx context.Context, resource *entity.Resource) error {
	_, err := r.col.InsertOne(ctx, resource)
	return err
}
