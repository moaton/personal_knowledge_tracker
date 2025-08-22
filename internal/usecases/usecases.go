package usecases

import (
	"context"
	"personal_knowledge_tracker/internal/interfaces"

	"github.com/go-logr/logr"
)

type Dependencies struct {
	Ctx    context.Context
	Repo   interfaces.Repository
	Logger logr.Logger
}

type usecases struct {
	user     interfaces.UserUsecases
	review   interfaces.ReviewUsecases
	resource interfaces.ResourceUsecases
}

func New(deps Dependencies) *usecases {
	return &usecases{
		user:     NewUserUsecases(deps),
		review:   NewReviewUsecase(deps),
		resource: NewResourceUsecases(deps),
	}
}

var _ interfaces.Usecases = (*usecases)(nil)

func (u *usecases) User() interfaces.UserUsecases {
	return u.user
}

func (u *usecases) Review() interfaces.ReviewUsecases {
	return u.review
}

func (u *usecases) Resource() interfaces.ResourceUsecases {
	return u.resource
}
