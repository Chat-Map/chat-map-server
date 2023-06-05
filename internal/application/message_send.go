package application

import (
	"context"
	"fmt"

	"github.com/Chat-Map/chat-map-server/internal/core"
)

type StoreMessageCommandRequest struct {
	ChatID  int64  `validate:"required" json:"chat_id"`
	UserID  int64  `validate:"required" json:"user_id"`
	Content string `validate:"required" json:"content"`
}

type StoreMessageCommand interface {
	Execute(ctx context.Context, params StoreMessageCommandRequest) (int64, error)
}

type StoreMessageCommandImplV1 struct {
	v  Validator
	ur UserRepository
	cr ChatRepository
	mr MessageRepository
	sr SessionsRepository
}

func NewStoreMessageCommandImplV1(
	v Validator,
	ur UserRepository,
	cr ChatRepository,
	mr MessageRepository,
	sr SessionsRepository,
) StoreMessageCommand {
	return StoreMessageCommandImplV1{v: v, ur: ur, cr: cr, mr: mr, sr: sr}
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
		return 0, fmt.Errorf("message sender id mismatch with payload")
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
