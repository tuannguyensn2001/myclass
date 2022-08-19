package userservice

import (
	"context"
	"myclass/src/models"
)

func (s *service) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.repository.FindByEmail(ctx, email)
}
