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

	result, err := db.NamedExec(`insert into users (email,username,password,created_at, updated_at) values (:email,:username,:password,:created_at,:updated_at)`, user)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.Id = int(id)

	return nil

}
