package application

import (
	"context"

	"github.com/Chat-Map/chat-map-server/internal/core"
)

type SignupCommandRequest struct {
	FirstName string `validate:"required,alpha"`
	LastName  string `validate:"required,alpha"`
	Phone     string `validate:"required,iso3166_1_alpha_numeric"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required,min=8"`
}

type SignupCommand interface {
	Execute(ctx context.Context, params SignupCommandRequest) error
}

type SignupCommandImplV1 struct {
	ur UserRepository
	ph PasswordHasher
	v  Validator
}

func NewSignupCommandImplV1(v Validator, ur UserRepository, ph PasswordHasher) SignupCommand {
	return SignupCommandImplV1{v: v, ur: ur, ph: ph}
}

func (s SignupCommandImplV1) Execute(ctx context.Context, params SignupCommandRequest) error {
	// Validate
	err := s.v.Validate(ctx, params)
	if err != nil {
		return err
	}
	// Hash password
	hashedPassword, err := s.ph.Hash(ctx, params.Password)
	if err != nil {
		return err
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
		return err
	}
	return nil
}