package usecases

import (
	"context"
	"personal_knowledge_tracker/internal/dto"
)

type user struct {
}

func NewUserUsecases() *user {
	return &user{}
}

func (u *user) Create(ctx context.Context, user *dto.User) error {
	return nil
}
