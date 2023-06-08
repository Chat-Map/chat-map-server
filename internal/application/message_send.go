package application

import (
	"context"

	"github.com/Chat-Map/chat-map-server/internal/core"
	"github.com/lordvidex/errs"
)

type StoreMessageCommandRequest struct {
	ChatID  int64  `validate:"required"`
	UserID  int64  `validate:"required"`
	Content string `validate:"required"`
}

type StoreMessageCommand interface {
	Execute(ctx context.Context, params StoreMessageCommandRequest) (int64, error)
}

type StoreMessageCommandImplV1 struct {
	v  Validator
	ur UserRepository
	cr ChatRepository
	mr MessageRepository
}

func NewStoreMessageCommandImplV1(
	v Validator,
	ur UserRepository,
	cr ChatRepository,
	mr MessageRepository,
) StoreMessageCommand {
	return StoreMessageCommandImplV1{v: v, ur: ur, cr: cr, mr: mr}
}

func (s StoreMessageCommandImplV1) Execute(ctx context.Context, params StoreMessageCommandRequest) (int64, error) {
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
	// Check that sender is the token user
	if payload.UserID != params.UserID {
		return 0, errs.B().Code(errs.Forbidden).Msg("message sender id mismatch with payload").Err()
	}
	// Check that user is member of the chat
	isMember, err := s.cr.IsChatMember(ctx, params.ChatID, params.UserID)
	if err != nil {
		return 0, err
	}
	if !isMember {
		return 0, errNotChatMember(params.ChatID, payload.UserID)
	}
	// Store message
	id, err := s.mr.StoreMessage(ctx, params.ChatID, core.Message{
		SenderID: params.UserID,
		Content:  params.Content,
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}
