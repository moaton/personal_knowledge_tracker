package mongo

import (
	"personal_knowledge_tracker/internal/interfaces"

	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	userRepository     *userRepository
	reviewRepository   *reviewRepository
	resourceRepository *resourceRepository
	fs                 *fs
}

func NewMongoRepository(db *mongo.Database) *mongoRepository {
	return &mongoRepository{
		userRepository:     NewUserRepository(db),
		reviewRepository:   NewReviewRepository(db),
		resourceRepository: NewResourceRepository(db),
		fs:                 NewFS(db),
	}
}

var _ interfaces.MongoRepository = (*mongoRepository)(nil)

func (r *mongoRepository) User() interfaces.UserRepository {
	return r.userRepository
}

func (r *mongoRepository) Review() interfaces.ReviewRepository {
	return r.reviewRepository
}

func (r *mongoRepository) Resource() interfaces.ResourceRepository {
	return r.resourceRepository
}

func (r *mongoRepository) FileStorage() interfaces.FileStorageRepository {
	return r.fs
}
