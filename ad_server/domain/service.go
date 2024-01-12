package domain

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, ad *Ad) error
	Get(ctx context.Context, id *IdAd) (*Ad, error)
	ServeAd(ctx context.Context, id *IdAd) (*AdServe, error)
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

func (s *Service) Create(ctx context.Context, ad *Ad) error {
	err := s.rp.Create(ctx, ad)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Get(ctx context.Context, id *IdAd) (*Ad, error) {
	ad, err := s.rp.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return ad, nil
}

func (s *Service) ServeAd(ctx context.Context, id *IdAd) (*AdServe, error) {
	url, err := s.rp.ServeAd(ctx, id)
	if err != nil {
		return nil, err
	}
	return url, nil
}
