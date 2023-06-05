package application

import (
	"context"

	"github.com/Chat-Map/chat-map-server/internal/core"
)

type TokenValidateCommandRequest struct {
	Token string `validate:"required"`
}

type TokenValidateCommand interface {
	Execute(ctx context.Context, params TokenValidateCommandRequest) (core.Payload, error)
}

type TokenValidateCommandImplV1 struct {
	v  Validator
	tk Tokenizer
}

func NewTokenValidateCommandImplV1(v Validator, tk Tokenizer) TokenValidateCommand {
	return &TokenValidateCommandImplV1{v: v, tk: tk}
}

func (s TokenValidateCommandImplV1) Execute(ctx context.Context, params TokenValidateCommandRequest) (core.Payload, error) {
	// Validate
	err := s.v.Validate(ctx, params)
	if err != nil {
		return core.Payload{}, err
	}
	// Validate token
	payload, err := s.tk.ValidateToken(ctx, params.Token)
	if err != nil {
		return core.Payload{}, err
	}
	return payload, nil
}
