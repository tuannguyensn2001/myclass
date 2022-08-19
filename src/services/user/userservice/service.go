package userservice

import (
	"context"
	"myclass/src/models"
)

type iRepository interface {
	Insert(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
}

type service struct {
	repository iRepository
}

func New(repository iRepository) *service {
	return &service{}
}
