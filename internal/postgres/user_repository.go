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

func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) (created *model.User, err error) {
	db := getTxOrDB(ctx, r.db).WithContext(ctx)

	userCreated := new(model.User)
	*userCreated = *user
	if err := db.Create(userCreated).Error; err != nil {
		return nil, err
	}

	return userCreated, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *model.User) (updated *model.User, err error) {
	db := getTxOrDB(ctx, r.db).WithContext(ctx)

	userUpdated := new(model.User)
	*userUpdated = *user
	if err := db.Save(userUpdated).Error; err != nil {
		return nil, err
	}

	return userUpdated, nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int64) error {
	db := getTxOrDB(ctx, r.db).WithContext(ctx)

	if err := db.Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUser(ctx context.Context, id int64) (*model.User, error) {
	db := getTxOrDB(ctx, r.db).WithContext(ctx)

	user := new(model.User)
	err := db.First(&user, id).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, pkgerr.CodeError{
			Code:      pkgerr.CodeNotFound,
			Err:       err,
			ClientMsg: fmt.Sprintf("user with id %d not found", id),
		}
	case err != nil:
		return nil, err
	}

	return user, nil
}
