package application

import (
	"context"
)

type CreateChatCommandRequest struct {
	UserID int64 `validate:"required" json:"user_id"`
}

type CreateChatCommand interface {
	Execute(ctx context.Context, params CreateChatCommandRequest) (int64, error)
}

type CreateChatCommandImplV1 struct {
	v  Validator
	cr ChatRepository
}

func NewCreateChatCommandImplV1(v Validator, cr ChatRepository) CreateChatCommand {
	return CreateChatCommandImplV1{v: v, cr: cr}
}

func (s CreateChatCommandImplV1) Execute(ctx context.Context, params CreateChatCommandRequest) (int64, error) {
	// Validate
	err := s.v.Validate(ctx, params)
	if err != nil {
		return 0, err
	}
	// Get Payload
	payload, err := GetPayload(ctx)
	if err != nil {
		return 0, err
	}
	// Create chat
	id, err := s.cr.CreatePrivateChat(ctx, []int64{params.UserID, payload.UserID})
	if err != nil {
		return 0, err
	}
	return id, nil
}
