package usecases

import (
	"context"
	"personal_knowledge_tracker/internal/dto"
	"personal_knowledge_tracker/internal/interfaces"
)

type review struct {
	repo interfaces.ReviewRepository
}

func NewReviewUsecase(deps Dependencies) *review {
	return &review{
		repo: deps.Repo.Mongo().Review(),
	}
}

func (r *review) Create(ctx context.Context, review *dto.Review) error {
	return nil
}
