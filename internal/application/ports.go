package application

import (
	"context"

	"github.com/Chat-Map/chat-map-server/internal/core"
	"github.com/google/uuid"
)

type UserRepository interface {
	StoreUser(ctx context.Context, user core.User) error
	GetUser(ctx context.Context, userID int32) (core.User, error)
	GetByEmail(ctx context.Context, email string) (core.User, error)
	GetAllUsers(ctx context.Context) ([]core.UserBySearch, error)
	SearchUserByEmail(ctx context.Context, email string) ([]core.UserBySearch, error)
}

type ChatRepository interface {
	GetChat(ctx context.Context, chatID int32) (core.Chat, error)
	CreatePrivateChat(ctx context.Context, userIDs []int32) error
	GetChatsMetadata(ctx context.Context, userID int32) ([]core.ChatMetaData, error)
	IsChatMember(ctx context.Context, chatID int32, userID int32) (bool, error)
}

type MessageRepository interface {
	StoreMessage(ctx context.Context, chatID int32, message core.Message) (int32, error)
}

type SessionsRepository interface {
	StoreSession(ctx context.Context, session core.Session) error
	GetSession(ctx context.Context, sessionID uuid.UUID) (core.Session, error)
}

type PasswordHasher interface {
	Hash(ctx context.Context, password string) (string, error)
	Compare(ctx context.Context, hash, password string) bool
}

type Tokenizer interface {
	GenerateToken(ctx context.Context, payload core.Payload) (string, error)
	ValidateToken(ctx context.Context, token string) (core.Payload, error)
}

type Validator interface {
	Validate(ctx context.Context, data interface{}) error
}

type Server interface {
	Run(ctx context.Context, port string) error
}
