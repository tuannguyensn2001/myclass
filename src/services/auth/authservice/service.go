package authservice

import (
	"context"
	"errors"
	"myclass/src/models"
)

type iRepository interface {
}

type service struct {
	repository  iRepository
	userService iUserService
}

type iUserService interface {
	Create(ctx context.Context, user *models.User) error
}

func New() *service {
	return &service{}
}

func (s *service) SetUserService(service iUserService) {
	s.userService = service
}

func (s *service) CheckHealth() error {
	if s.userService == nil {
		return errors.New("user servive not defined in auth")
	}
	return nil
}
