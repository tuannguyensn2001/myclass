package userrepository

import (
	"context"
	baserepository "myclass/src/base/repository"
	"myclass/src/models"
)

func (r *repository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	db := baserepository.GetDBFromContext(ctx)

	var result models.User

	err := db.Select("id", "username", "email", "password", "created_at", "updated_at").Where("email = ?", email).First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil

}
