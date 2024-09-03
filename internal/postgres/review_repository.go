package postgres

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/isutare412/goasis/internal/core/model"
	"github.com/isutare412/goasis/internal/pkgerr"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(client *Client) *ReviewRepository {
	return &ReviewRepository{
		db: client.db,
	}
}

func (r *ReviewRepository) CreateReview(ctx context.Context, review *model.Review) (created *model.Review, err error) {
	db := getTxOrDB(ctx, r.db).WithContext(ctx)

	if err := db.Create(review).Error; err != nil {
		return nil, err
	}

	return review, nil
}

func (r *ReviewRepository) UpdateReview(ctx context.Context, review *model.Review) (updated *model.Review, err error) {
	db := getTxOrDB(ctx, r.db).WithContext(ctx)

	if err := db.Save(review).Error; err != nil {
		return nil, err
	}

	return review, nil
}

func (r *ReviewRepository) DeleteReview(ctx context.Context, id int64) error {
	db := getTxOrDB(ctx, r.db).WithContext(ctx)

	if err := db.Delete(&model.Review{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *ReviewRepository) GetReview(ctx context.Context, id int64) (*model.Review, error) {
	db := getTxOrDB(ctx, r.db).WithContext(ctx)

	var review model.Review
	err := db.
		First(&review, id).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, pkgerr.CodeError{
			Code:      pkgerr.CodeNotFound,
			Err:       err,
			ClientMsg: fmt.Sprintf("review with id %d not found", id),
		}
	case err != nil:
		return nil, err
	}

	return &review, nil
}

func (r *ReviewRepository) GetReviewPreload(ctx context.Context, id int64) (*model.Review, error) {
	db := getTxOrDB(ctx, r.db).WithContext(ctx)

	var review model.Review
	err := db.
		Joins("User").
		First(&review, id).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, pkgerr.CodeError{
			Code:      pkgerr.CodeNotFound,
			Err:       err,
			ClientMsg: fmt.Sprintf("review with id %d not found", id),
		}
	case err != nil:
		return nil, err
	}

	return &review, nil
}

func (r *ReviewRepository) ListReviewsOfCafe(ctx context.Context, cafeID int64) ([]*model.Review, error) {
	db := getTxOrDB(ctx, r.db).WithContext(ctx)

	var reviews []*model.Review
	err := db.
		Where(&model.Review{
			CafeID: cafeID,
		}).
		Order("updated_at desc").
		Find(&reviews).Error
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *ReviewRepository) ListReviewsOfUser(ctx context.Context, userID int64) ([]*model.Review, error) {
	db := getTxOrDB(ctx, r.db).WithContext(ctx)

	var reviews []*model.Review
	err := db.
		Where(&model.Review{
			UserID: userID,
		}).
		Order("updated_at desc").
		Find(&reviews).Error
	if err != nil {
		return nil, err
	}

	return reviews, nil
}
