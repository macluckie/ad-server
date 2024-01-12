package domain

import (
	"context"
)

type Repository interface {
	GetCount(ctx context.Context, id *IdAd) (*CountAd, error)
	IncrementCount(ctx context.Context, id *IdAd) error
}

type Service struct {
	rp Repository
}

func NewService(repo Repository) *Service {
	service := Service{
		rp: repo,
	}
	return &service
}

func (s *Service) GetCount(ctx context.Context, id *IdAd) (*CountAd, error) {
	count, err := s.rp.GetCount(ctx, id)
	if err != nil {
		return nil, err
	}
	return count, nil
}

func (s *Service) IncrementCount(ctx context.Context, id *IdAd) error {
	err := s.rp.IncrementCount(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
