package interfaces

import (
	"context"
	"personal_knowledge_tracker/internal/dto"
)

type (
	Usecases interface {
		User() UserUsecases
		Review() ReviewUsecases
		Resource() ResourceUsecases
	}
	UserUsecases interface {
		Create(ctx context.Context, user *dto.User) error
	}
	ReviewUsecases interface {
		Create(ctx context.Context, review *dto.Review) error
	}
	ResourceUsecases interface {
		Create(ctx context.Context, resource *dto.Resource) error
	}
)
