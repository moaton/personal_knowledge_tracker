package usecases

import (
	"context"
	"personal_knowledge_tracker/internal/dto"
	"personal_knowledge_tracker/internal/interfaces"
)

type user struct {
	repo interfaces.UserRepository
}

func NewUserUsecases(deps Dependencies) *user {
	return &user{
		repo: deps.Repo.Mongo().User(),
	}
}

func (u *user) Create(ctx context.Context, user *dto.User) error {
	return nil
}
