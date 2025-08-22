package interfaces

import (
	"context"
	"personal_knowledge_tracker/internal/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Repository interface {
		Mongo() MongoRepository
	}
	MongoRepository interface {
		User() UserRepository
		Review() ReviewRepository
		Resource() ResourceRepository
		FileStorage() FileStorageRepository
	}
	UserRepository interface {
		Create(ctx context.Context, user *entity.User) error
	}
	ReviewRepository interface {
		Create(ctx context.Context, review *entity.Review) error
	}
	ResourceRepository interface {
		Create(ctx context.Context, resource *entity.Resource) error
		List(ctx context.Context, userID int64, page, limit int64) ([]*entity.Resource, int64, error)
	}
	FileStorageRepository interface {
		Save(ctx context.Context, file *entity.File, data []byte) error
		Get(ctx context.Context, id primitive.ObjectID) ([]byte, error)
		Delete(ctx context.Context, id primitive.ObjectID) error
	}
)
