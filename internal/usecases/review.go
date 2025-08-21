package usecases

import (
	"context"
	"personal_knowledge_tracker/internal/dto"
)

type review struct {
}

func NewReviewUsecase() *review {
	return &review{}
}

func (r *review) Create(ctx context.Context, review *dto.Review) error {
	return nil
}
