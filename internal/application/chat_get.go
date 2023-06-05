package application

import (
	"context"

	"github.com/Chat-Map/chat-map-server/internal/core"
)

type GetChatCommandRequest struct {
	ChatID int32 `validate:"required" json:"chat_id"`
}

type GetChatCommand interface {
	Execute(ctx context.Context, params GetChatCommandRequest) (core.Chat, error)
}

type GetChatCommandImplV1 struct {
	v  Validator
	cr ChatRepository
}

func NewGetChatCommandImplV1(v Validator, cr ChatRepository) GetChatCommand {
	return GetChatCommandImplV1{v: v, cr: cr}
}

func (s GetChatCommandImplV1) Execute(ctx context.Context, params GetChatCommandRequest) (core.Chat, error) {
	// Validate
	err := s.v.Validate(ctx, params)
	if err != nil {
		return core.Chat{}, err
	}
	// Get Payload
	payload, err := GetPayload(ctx)
	if err != nil {
		return core.Chat{}, err
	}
	// Check that user is member of the chat
	isMember, err := s.cr.IsChatMember(ctx, params.ChatID, payload.UserID)
	if err != nil {
		return core.Chat{}, err
	}
	if !isMember {
		return core.Chat{}, errNotChatMember(params.ChatID, payload.UserID)
	}
	// Get chat
	chat, err := s.cr.GetChat(ctx, params.ChatID)
	if err != nil {
		return core.Chat{}, err
	}
	return chat, nil
}
