package application

import (
	"context"

	"github.com/Chat-Map/chat-map-server/internal/core"
	"github.com/google/uuid"
)

type UserRepository interface {
	StoreUser(ctx context.Context, user core.User) error
	GetUser(ctx context.Context, userID int64) (core.User, error)
	GetByEmail(ctx context.Context, email string) (core.User, error)
	GetAllUsers(ctx context.Context) ([]core.UserBySearch, error)
	SearchUserByAll(ctx context.Context, pattern string) ([]core.UserBySearch, error)
}

type ChatRepository interface {
	GetChat(ctx context.Context, chatID int64) (core.Chat, error)
	CreatePrivateChat(ctx context.Context, userIDs []int64) (int64, error)
	GetChatsMetadata(ctx context.Context, userID int64) ([]core.ChatMetaData, error)
	IsChatMember(ctx context.Context, chatID int64, userID int64) (bool, error)
}

type MessageRepository interface {
	StoreMessage(ctx context.Context, chatID int64, message core.Message) (int64, error)
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
	Run(port string) error
}
