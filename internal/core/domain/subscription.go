package domain

import (
	"context"
	"time"
)

type Subscription struct {
	ID          string
	UserID      string
	TariffID    string
	TrafficUsed uint64
	StartAt     time.Time
	ExpiresAt   time.Time
}

func (s Subscription) IsActive(now time.Time) (isActive bool) {
	return now.Before(s.ExpiresAt)
}

func (s Subscription) IsExpired(now time.Time) (isExpired bool) {
	return now.After(s.ExpiresAt)
}

func (s *Subscription) AddTraffic(bytes uint64) {
	s.TrafficUsed += bytes
}

type SubscriptionRepository interface {
	Create(ctx context.Context, sub Subscription) (Subscription, error)
	Update(ctx context.Context, sub Subscription) (Subscription, error)
	Delete(ctx context.Context, id string) error

	GetByID(ctx context.Context, id string) (Subscription, error)
	GetByUserID(ctx context.Context, userID string) ([]Subscription, error)
}
