package mongo

import (
	"context"
	"personal_knowledge_tracker/internal/entity"

	"go.mongodb.org/mongo-driver/mongo"
)

type reviewRepository struct {
	col *mongo.Collection
}

func NewReviewRepository(db *mongo.Database) *reviewRepository {
	return &reviewRepository{col: db.Collection("reviews")}
}

func (r *reviewRepository) Create(ctx context.Context, review *entity.Review) error {
	_, err := r.col.InsertOne(ctx, review)
	return err
}
