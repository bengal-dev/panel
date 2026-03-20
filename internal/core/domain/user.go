package domain

import (
	"context"
	"time"
)

type UserConfig struct {
	Protocol TunnelProtocol
}

type User struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	LastLogin time.Time
	IsBlocked bool
	BlockedAt *time.Time

	Config UserConfig
}

func (u *User) Block(now time.Time) {
	u.IsBlocked = true
	u.BlockedAt = &now
}

func (u *User) Unblock() {
	u.IsBlocked = false
	u.BlockedAt = nil
}

func (u *User) TouchLogin(now time.Time) {
	u.LastLogin = now
}

type UserRepository interface {
	Create(ctx context.Context, user User) (newUser User, err error)
	Update(ctx context.Context, user User) (updatedUser User, err error)

	GetByID(ctx context.Context, id string) (user User, err error)
}

type UserService interface {
	BlockUser(ctx context.Context, id string) (err error)
	UnblockUser(ctx context.Context, id string) (err error)
	UpdateLastLogin(ctx context.Context, id string) (err error)
}
