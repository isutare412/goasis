package cafe

import (
	"context"
	"fmt"

	"github.com/isutare412/goasis/internal/core/model"
	"github.com/isutare412/goasis/internal/core/port"
)

type Service struct {
	cafeRepository port.CafeRepository
}

func NewService(cafeRepository port.CafeRepository) *Service {
	return &Service{
		cafeRepository: cafeRepository,
	}
}

func (s *Service) CreateCafe(ctx context.Context, cafe model.Cafe) (created model.Cafe, err error) {
	cafeCreated, err := s.cafeRepository.CreateCafe(ctx, cafe)
	if err != nil {
		return model.Cafe{}, fmt.Errorf("creating cafe: %w", err)
	}

	return cafeCreated, nil
}

func (s *Service) UpdateCafe(ctx context.Context, cafe model.Cafe) (updated model.Cafe, err error) {
	cafeUpdated, err := s.cafeRepository.UpdateCafe(ctx, cafe)
	if err != nil {
		return model.Cafe{}, fmt.Errorf("updating cafe: %w", err)
	}

	return cafeUpdated, nil
}

func (s *Service) DeleteCafe(ctx context.Context, cafeID int64) error {
	if err := s.cafeRepository.DeleteCafe(ctx, cafeID); err != nil {
		return fmt.Errorf("deleting cafe: %w", err)
	}

	return nil
}

func (s *Service) GetCafe(ctx context.Context, cafeID int64) (model.Cafe, error) {
	cafe, err := s.cafeRepository.GetCafe(ctx, cafeID)
	if err != nil {
		return model.Cafe{}, fmt.Errorf("getting cafe: %w", err)
	}

	return cafe, nil
}

func (s *Service) ListCafes(ctx context.Context) ([]model.Cafe, error) {
	cafes, err := s.cafeRepository.ListCafes(ctx)
	if err != nil {
		return nil, fmt.Errorf("listing cafes: %w", err)
	}

	return cafes, nil
}
