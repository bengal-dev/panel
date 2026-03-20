package domain

import (
	"context"
	"errors"
	"time"
)

type AuthProviderType string

const (
	TelegramProvider AuthProviderType = "telegram"
	DiscordProvider  AuthProviderType = "discord"
	EmailProvider    AuthProviderType = "email"
	PasswordProvider AuthProviderType = "password"
)

type TokenType string

const (
	TokenAccess  TokenType = "access"
	TokenRefresh TokenType = "refresh"
)

type MFAType string

const (
	TOTP MFAType = "totp"
)

var (
	ErrUserNotFound             = errors.New("user not found")
	ErrUserBlocked              = errors.New("user is blocked")
	ErrIdentityAlreadyExists    = errors.New("identity already exists")
	ErrMFAInvalidVerification   = errors.New("mfa invalid verification")
	ErrMFANotCompletedChallenge = errors.New("mfa not completed challenge")
)

type MFA struct {
	UserID    string
	Type      MFAType
	Secret    string
	Activated bool
}

type Credential struct {
	UserID       string
	PasswordHash string
}

type AuthIdentity struct {
	Type  AuthProviderType
	Value string
}

type Identity struct {
	ID     string
	UserID string
	AuthIdentity
}

type Token struct {
	Value     string
	Type      TokenType
	ExpiresAt time.Time
}

type TokenPair struct {
	Access  Token
	Refresh Token
}

type Session struct {
	ID        string
	UserID    string
	TokenHash string // refresh token
	ExpiresAt time.Time
}

type UserProvider interface {
	GetByID(ctx context.Context, id string) (user *User, err error)
}

type IdentityRepository interface {
	Get(ctx context.Context, authIdentity AuthIdentity) (identity *Identity, err error)
	AddIdentity(ctx context.Context, identity *Identity) (err error)
}

type MFARepository interface {
	GetByUserID(ctx context.Context, userID string) (mfa *MFA, err error)
	Create(ctx context.Context, mfa *MFA) (err error)
	Delete(ctx context.Context, userID string) (err error)
	Update(ctx context.Context, mfa *MFA) (err error)
}

type MFAService interface {
	VerifyMFA(ctx context.Context, userID string, code string) (err error)
}

type AuthService interface {
	Login(ctx context.Context, identity AuthIdentity) (tokens *TokenPair, err error)
	Logout(ctx context.Context, sessionID string) (err error)
	LogoutAll(ctx context.Context, userID string) (err error)
	Refresh(ctx context.Context, refreshToken string) (tokens *TokenPair, err error)
}

type SessionRepository interface {
	Create(ctx context.Context, session *Session) (err error)
	GetByHash(ctx context.Context, hash string) (session *Session, err error)
	DeleteByUserID(ctx context.Context, userID string) (err error)
	DeleteByID(ctx context.Context, id string) (err error)
}

type CredentialRepository interface {
	GetByUserID(ctx context.Context, userID string) (cred *Credential, err error)
	Create(ctx context.Context, cred *Credential) (err error)
	Update(ctx context.Context, cred *Credential) (err error)
}
