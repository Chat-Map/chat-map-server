package application

import (
	"context"
)

type CreateChatCommandRequest struct {
	UserID int64 `validate:"required"`
}

type CreateChatCommandResponse struct {
	ChatID int64
}

type CreateChatCommand interface {
	Execute(ctx context.Context, params CreateChatCommandRequest) (CreateChatCommandResponse, error)
}

type CreateChatCommandImplV1 struct {
	v  Validator
	cr ChatRepository
	cn ChatNotifier
}

func NewCreateChatCommandImplV1(v Validator, cr ChatRepository, cn ChatNotifier) CreateChatCommand {
	return CreateChatCommandImplV1{v: v, cr: cr, cn: cn}
}

func (s CreateChatCommandImplV1) Execute(ctx context.Context, params CreateChatCommandRequest) (CreateChatCommandResponse, error) {
	// Validate
	err := s.v.Validate(ctx, params)
	if err != nil {
		return CreateChatCommandResponse{}, err
	}
	// Get Payload
	payload, err := GetPayload(ctx)
	if err != nil {
		return CreateChatCommandResponse{}, err
	}
	// Create chat
	userIDs := []int64{params.UserID, payload.UserID}
	chatID, err := s.cr.CreatePrivateChat(ctx, userIDs)
	if err != nil {
		return CreateChatCommandResponse{}, err
	}
	// Notify users about newly created chat
	go func() {
		s.cn.Notify(ctx, userIDs, chatID)
	}()
	return CreateChatCommandResponse{ChatID: chatID}, nil
}
