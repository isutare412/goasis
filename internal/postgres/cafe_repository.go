package postgres

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/isutare412/goasis/internal/core/model"
	"github.com/isutare412/goasis/internal/pkgerr"
)

type CafeRepository struct {
	db *gorm.DB
}

func NewCafeRepository(client *Client) *CafeRepository {
	return &CafeRepository{
		db: client.db,
	}
}

func (r *CafeRepository) CreateCafe(ctx context.Context, cafe *model.Cafe) (created *model.Cafe, err error) {
	db := getTxOrDB(ctx, r.db).WithContext(ctx)

	if err := db.Create(cafe).Error; err != nil {
		return nil, err
	}

	return cafe, nil
}

func (r *CafeRepository) UpdateCafe(ctx context.Context, cafe *model.Cafe) (updated *model.Cafe, err error) {
	db := getTxOrDB(ctx, r.db).WithContext(ctx)

	if err := db.Save(cafe).Error; err != nil {
		return nil, err
	}

	return cafe, nil
}

func (r *CafeRepository) DeleteCafe(ctx context.Context, id int64) error {
	db := getTxOrDB(ctx, r.db).WithContext(ctx)

	if err := db.Delete(&model.Cafe{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *CafeRepository) GetCafe(ctx context.Context, id int64) (*model.Cafe, error) {
	db := getTxOrDB(ctx, r.db).WithContext(ctx)

	var cafe model.Cafe
	err := db.First(&cafe, id).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, pkgerr.CodeError{
			Code:      pkgerr.CodeNotFound,
			Err:       err,
			ClientMsg: fmt.Sprintf("cafe with id %d not found", id),
		}
	case err != nil:
		return nil, err
	}

	return &cafe, nil
}

func (r *CafeRepository) ListCafes(ctx context.Context) ([]*model.Cafe, error) {
	db := getTxOrDB(ctx, r.db).WithContext(ctx)

	var cafes []*model.Cafe
	err := db.
		Order("id").
		Find(&cafes).Error
	if err != nil {
		return nil, err
	}

	return cafes, nil
}
