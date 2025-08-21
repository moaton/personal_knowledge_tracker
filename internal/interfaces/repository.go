package interfaces

import (
	"context"
	"personal_knowledge_tracker/internal/entity"
)

type (
	Repository interface {
		Mongo() MongoRepository
	}
	MongoRepository interface {
		User() UserRepository
		Review() ReviewRepository
		Resource() ResourceRepository
	}
	UserRepository interface {
		Create(ctx context.Context, user *entity.User) error
	}
	ReviewRepository interface {
		Create(ctx context.Context, review *entity.Review) error
	}
	ResourceRepository interface {
		Create(ctx context.Context, resource *entity.Resource) error
	}
)
