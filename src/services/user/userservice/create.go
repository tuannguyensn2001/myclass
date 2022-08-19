package userservice

import (
	"context"
	"myclass/src/models"
)

func (s *service) Create(ctx context.Context, user *models.User) error {
	return s.repository.Insert(ctx, user)
}
