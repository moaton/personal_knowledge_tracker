package repository

import (
	"personal_knowledge_tracker/internal/interfaces"
	mongoRepo "personal_knowledge_tracker/internal/repository/mongo"

	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	mongoRepo interfaces.MongoRepository
}

func New(db *mongo.Database) *repository {
	return &repository{
		mongoRepo: mongoRepo.NewMongoRepository(db),
	}
}

var _ interfaces.Repository = (*repository)(nil)

func (r *repository) Mongo() interfaces.MongoRepository {
	return r.mongoRepo
}
