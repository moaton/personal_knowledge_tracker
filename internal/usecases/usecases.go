package usecases

import (
	"context"

	"github.com/go-logr/logr"
)

type Dependencies struct {
	Ctx    context.Context
	Logger logr.Logger
}

type usecases struct {
}

func New(deps Dependencies) *usecases {
	return &usecases{}
}
