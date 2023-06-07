package application

import (
	"context"
	"time"

	"github.com/Chat-Map/chat-map-server/internal/core"
)

type SignupCommandRequest struct {
	FirstName string `validate:"required,alpha"`
	LastName  string `validate:"required,alpha"`
	Phone     string `validate:"required,e164"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required,min=8"`
}

type SignupCommandResponse struct {
	User         core.User
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

type SignupCommand interface {
	Execute(ctx context.Context, params SignupCommandRequest) (SignupCommandResponse, error)
}

type SignupCommandImplV1 struct {
	ur UserRepository
	ph PasswordHasher
	v  Validator
	in SigninCommand
}

func NewSignupCommandImplV1(v Validator, ur UserRepository, ph PasswordHasher, in SigninCommand) SignupCommand {
	return SignupCommandImplV1{v: v, ur: ur, ph: ph, in: in}
}

func (s SignupCommandImplV1) Execute(ctx context.Context, params SignupCommandRequest) (SignupCommandResponse, error) {
	// Validate
	err := s.v.Validate(ctx, params)
	if err != nil {
		return SignupCommandResponse{}, err
	}
	// Hash password
	hashedPassword, err := s.ph.Hash(ctx, params.Password)
	if err != nil {
		return SignupCommandResponse{}, err
	}
	// Store user
	err = s.ur.StoreUser(ctx, core.User{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Phone:     params.Phone,
		Email:     params.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		return SignupCommandResponse{}, err
	}
	response, err := s.in.Execute(ctx, SigninCommandRequest{
		Email:    params.Email,
		Password: params.Password,
	})
	return SignupCommandResponse{
		User:         response.User,
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
		ExpiresAt:    response.ExpiresAt,
	}, nil
}
