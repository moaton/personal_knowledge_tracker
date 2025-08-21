package usecases

import (
	"context"
	"personal_knowledge_tracker/internal/dto"
)

type resource struct {
}

func NewResourceUsecases() *resource {
	return &resource{}
}

func (r *resource) Create(ctx context.Context, resource *dto.Resource) error {
	return nil
}
