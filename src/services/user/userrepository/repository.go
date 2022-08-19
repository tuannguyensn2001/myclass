package userrepository

import (
	"context"
	baserepository "myclass/src/base/repository"
	"myclass/src/models"
)

type repository struct {
}

func New() *repository {
	return &repository{}
}

func (r *repository) Insert(ctx context.Context, user *models.User) error {
	db := baserepository.GetDBFromContext(ctx)

	err := db.Create(user).Error
	if err != nil {
		return err
	}

	return nil

}
