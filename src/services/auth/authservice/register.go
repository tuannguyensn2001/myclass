package authservice

import (
	"context"
	"errors"
	"myclass/src/app"
	"myclass/src/models"
	"myclass/src/services/auth/authstruct"
	"myclass/src/services/auth/authutil"
)

func (s *service) Register(ctx context.Context, input authstruct.RegisterInput) error {

	check := s.userService.CheckUserExistedByEmail(ctx, input.Email)

	if check {
		return app.ConflictHttpError("user existed", errors.New("user existed"))
	}

	password, err := authutil.Hash(input.Password, 5)
	if err != nil {
		return err
	}

	user := models.User{
		Email:    input.Email,
		Username: input.Username,
		Password: password,
	}

	err = s.userService.Create(ctx, &user)
	if err != nil {
		return err
	}

	return nil

}
