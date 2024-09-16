package postgres

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/isutare412/goasis/internal/core/model"
	"github.com/isutare412/goasis/internal/pkgerr"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(client *Client) *UserRepository {
	return &UserRepository{
		db: client.db,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user model.User) (created model.User, err error) {
	db := getTxOrDB(ctx, r.db).WithContext(ctx)

	if err := db.Create(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user model.User) (updated model.User, err error) {
	db := getTxOrDB(ctx, r.db).WithContext(ctx)

	if err := db.Save(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int64) error {
	db := getTxOrDB(ctx, r.db).WithContext(ctx)

	if err := db.Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUser(ctx context.Context, id int64) (model.User, error) {
	db := getTxOrDB(ctx, r.db).WithContext(ctx)

	var user model.User
	err := db.First(&user, id).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return model.User{}, pkgerr.CodeError{
			Code:      pkgerr.CodeNotFound,
			Err:       err,
			ClientMsg: fmt.Sprintf("user with id %d not found", id),
		}
	case err != nil:
		return model.User{}, err
	}

	return user, nil
}
