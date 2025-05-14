package service

import (
	"context"
	"weather_forecast_sub/internal/domain"
	"weather_forecast_sub/internal/repository"
)

type SubscriptionService struct {
	repo repository.SubscriptionRepository
}

func NewSubscriptionService(repo repository.SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

func (s *SubscriptionService) Create(ctx context.Context, subscription domain.Subscription) error {
	return nil
}

func (s *SubscriptionService) Update(ctx context.Context, inp UpdateSubscriptionInput) error {
	return nil
}

func (s *SubscriptionService) Delete(ctx context.Context, token string) error {
	return nil
}
