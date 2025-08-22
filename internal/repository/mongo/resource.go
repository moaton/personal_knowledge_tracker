package mongo

import (
	"context"
	"personal_knowledge_tracker/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *resourceRepository) List(ctx context.Context, userID, page, limit int64) ([]*entity.Resource, int64, error) {
	skip := (page - 1) * limit

	findOptions := options.Find().
		SetSkip(skip).
		SetLimit(limit).
		SetSort(bson.D{{Key: "createdAt", Value: -1}})

	filter := bson.M{
		"metadata.userID": userID,
	}

	total, err := r.col.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	cursor, err := r.col.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var resources []*entity.Resource
	if err := cursor.All(ctx, &resources); err != nil {
		return nil, 0, err
	}

	return resources, total, nil
}
