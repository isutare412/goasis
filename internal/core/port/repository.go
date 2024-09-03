package port

import (
	"context"

	"github.com/isutare412/goasis/internal/core/model"
)

type UserRepository interface {
	CreateUser(context.Context, *model.User) (created *model.User, err error)
	UpdateUser(context.Context, *model.User) (updated *model.User, err error)
	DeleteUser(ctx context.Context, id int64) error
	GetUser(ctx context.Context, id int64) (*model.User, error)
}

type CafeRepository interface {
	CreateCafe(context.Context, *model.Cafe) (created *model.Cafe, err error)
	UpdateCafe(context.Context, *model.Cafe) (updated *model.Cafe, err error)
	DeleteCafe(ctx context.Context, id int64) error
	GetCafe(ctx context.Context, id int64) (*model.Cafe, error)
	ListCafes(context.Context) ([]*model.Cafe, error)
}

type ReviewRepository interface {
	CreateReview(context.Context, *model.Review) (created *model.Review, err error)
	UpdateReview(context.Context, *model.Review) (updated *model.Review, err error)
	DeleteReview(ctx context.Context, id int64) error
	GetReview(ctx context.Context, id int64) (*model.Review, error)
	GetReviewPreload(ctx context.Context, id int64) (*model.Review, error)
	ListReviewsOfCafe(ctx context.Context, cafeID int64) ([]*model.Review, error)
	ListReviewsOfUser(ctx context.Context, userID int64) ([]*model.Review, error)
}
