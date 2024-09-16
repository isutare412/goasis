package port

import (
	"context"

	"github.com/isutare412/goasis/internal/core/model"
)

type CafeService interface {
	CreateCafe(context.Context, model.Cafe) (created model.Cafe, err error)
	UpdateCafe(context.Context, model.Cafe) (updated model.Cafe, err error)
	DeleteCafe(ctx context.Context, cafeID int64) error
	GetCafe(ctx context.Context, cafeID int64) (model.Cafe, error)
	ListCafes(context.Context) ([]model.Cafe, error)
}
