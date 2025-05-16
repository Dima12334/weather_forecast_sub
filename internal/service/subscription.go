package service

import (
	"context"
	"time"
	"weather_forecast_sub/internal/domain"
	"weather_forecast_sub/internal/repository"
	"weather_forecast_sub/pkg/hash"
)

type SubscriptionService struct {
	repo   repository.SubscriptionRepository
	hasher hash.EmailHasher
}

func NewSubscriptionService(repo repository.SubscriptionRepository, hasher hash.EmailHasher) *SubscriptionService {
	return &SubscriptionService{repo: repo, hasher: hasher}
}

func (s *SubscriptionService) Create(ctx context.Context, inp CreateSubscriptionInput) error {
	token := s.hasher.GenerateEmailHash(inp.Email)

	subscription := domain.Subscription{
		CreatedAt:  time.Now(),
		Email:      inp.Email,
		City:       inp.City,
		Frequency:  inp.Frequency,
		Token:      token,
		Confirmed:  false,
		LastSentAt: nil,
	}
	err := s.repo.Create(ctx, subscription)

	// TODO: send confirmation email
	return err
}

func (s *SubscriptionService) Confirm(ctx context.Context, token string) error {
	_, err := s.repo.GetByToken(ctx, token)
	if err != nil {
		return err
	}

	return s.repo.Confirm(ctx, token)
}

func (s *SubscriptionService) Delete(ctx context.Context, token string) error {
	_, err := s.repo.GetByToken(ctx, token)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, token)
}
