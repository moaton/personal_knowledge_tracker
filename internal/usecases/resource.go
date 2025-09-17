package usecases

import (
	"context"
	"personal_knowledge_tracker/internal/dto"
	"personal_knowledge_tracker/internal/interfaces"
)

type resource struct {
	repo interfaces.ResourceRepository
}

func NewResourceUsecases(deps Dependencies) *resource {
	return &resource{
		repo: deps.Repo.Mongo().Resource(),
	}
}

func (r *resource) Create(ctx context.Context, resource *dto.Resource) error {
	return r.repo.Create(ctx, convertResourceToEntity(resource))
}

func (r *resource) List(ctx context.Context, userID, page, limit int64) ([]*dto.Resource, int64, error) {
	resources, total, err := r.repo.List(ctx, userID, page, limit)
	if err != nil {
		return nil, 0, err
	}

	return convertResourcesToDTO(resources), total, nil
}

func (r *resource) DeleteByID(ctx context.Context, id string) error {
	return r.repo.DeleteByID(ctx, id)
}
