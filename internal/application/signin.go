package application

import (
	"context"
	"fmt"
	"time"

	"github.com/Chat-Map/chat-map-server/internal/core"
	"github.com/google/uuid"
)

type SigninCommandRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

var (
	sessionLifetime      = time.Hour * 24 * 7 // 7 days
	refreshTokenLifetime = time.Hour * 24     // 1 day
	accessTokenLifetime  = time.Hour          // 1 hour
)

type SigninCommandResponse struct {
	User         core.User
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

type SigninCommand interface {
	Execute(ctx context.Context, params SigninCommandRequest) (SigninCommandResponse, error)
}

type SigninCommandImplV1 struct {
	v  Validator
	ur UserRepository
	sr SessionsRepository
	ph PasswordHasher
	tk Tokenizer
}

func NewSigninCommandImplV1(v Validator, ur UserRepository, sr SessionsRepository, ph PasswordHasher, tk Tokenizer) SigninCommand {
	return SigninCommandImplV1{v: v, ur: ur, sr: sr, ph: ph, tk: tk}
}

func (s SigninCommandImplV1) Execute(ctx context.Context, params SigninCommandRequest) (SigninCommandResponse, error) {
	// Vadidate
	err := s.v.Validate(ctx, params)
	if err != nil {
		return SigninCommandResponse{}, err
	}
	// Get user
	user, err := s.ur.GetByEmail(ctx, params.Email)
	if err != nil {
		return SigninCommandResponse{}, err
	}
	// Compare passwords
	similar := s.ph.Compare(ctx, user.Password, params.Password)
	if !similar {
		return SigninCommandResponse{}, fmt.Errorf("incorrect password")
	}
	// Create session
	sessionID := uuid.New()
	err = s.sr.StoreSession(ctx, core.Session{
		ID:        sessionID,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(sessionLifetime),
	})
	if err != nil {
		return SigninCommandResponse{}, err
	}
	// Generate tokens
	accessToken, err := s.tk.GenerateToken(ctx, core.Payload{
		UserID:    user.ID,
		SessionID: sessionID,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(accessTokenLifetime),
	})
	if err != nil {
		return SigninCommandResponse{}, err
	}
	refreshToken, err := s.tk.GenerateToken(ctx, core.Payload{
		UserID:    user.ID,
		SessionID: sessionID,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(refreshTokenLifetime),
	})
	if err != nil {
		return SigninCommandResponse{}, err
	}

	return SigninCommandResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(accessTokenLifetime),
	}, nil
}
