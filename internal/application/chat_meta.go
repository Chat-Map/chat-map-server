package application

import (
	"context"

	"github.com/Chat-Map/chat-map-server/internal/core"
)

type GetChatMetaCommandRequest struct {
}

type GetChatMetaCommandResponse struct {
	ChatMetaData []core.ChatMetaData
}

type GetChatMetaCommand interface {
	Execute(ctx context.Context, params GetChatMetaCommandRequest) (GetChatMetaCommandResponse, error)
}

type GetChatMetaCommandImplV1 struct {
	cr ChatRepository
}

func NewGetChatMetaCommandImplV1(cr ChatRepository) GetChatMetaCommand {
	return GetChatMetaCommandImplV1{cr: cr}
}

func (s GetChatMetaCommandImplV1) Execute(ctx context.Context, params GetChatMetaCommandRequest) (GetChatMetaCommandResponse, error) {
	// Get Payload
	payload, err := GetPayload(ctx)
	if err != nil {
		return GetChatMetaCommandResponse{}, err
	}
	// Get user's metadata
	metadata, err := s.cr.GetChatsMetadata(ctx, payload.UserID)
	if err != nil {
		return GetChatMetaCommandResponse{}, err
	}
	return GetChatMetaCommandResponse{ChatMetaData: metadata}, nil
}
