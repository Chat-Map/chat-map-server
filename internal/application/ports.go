package application

import (
	"context"

	"github.com/Chat-Map/chat-map-server/internal/core"
)

type UserRepository interface {
	StoreUser(ctx context.Context, user core.User) error
	GetUser(ctx context.Context, userID string) (core.User, error)
	SearchUser(ctx context.Context, userID string) ([]core.UserBySearch, error)
}

type ChatRepository interface {
	CreateChat(ctx context.Context, userID1 string, userID2 string) error
	GetChat(ctx context.Context, chatID string) (core.Chat, error)
	GetChatsMetadata(ctx context.Context, usersID string) ([]core.ChatMetaData, error)
}

type MessageRepository interface {
	StoreMessage(ctx context.Context, chatID string, message core.Message) error
	GetMessage(ctx context.Context, chatID string, messageID int) (core.Message, error)
	GetMessages(ctx context.Context, chatID string) ([]core.Message, error)
}

type SessionsRepository interface {
	StoreSession(ctx context.Context, session core.Session) error
	GetSession(ctx context.Context, sessionID string) (core.Session, error)
	DeleteSession(ctx context.Context, sessionID string) error
}

type Server interface {
	Run(ctx context.Context, port string) error
}
