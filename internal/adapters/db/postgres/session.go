package postgres

import (
	"context"
	"database/sql"

	"github.com/Chat-Map/chat-map-server/internal/core"
)

type SessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

// GetSession implements application.SessionsRepository
func (*SessionRepository) GetSession(ctx context.Context, sessionID string) (core.Session, error) {
	panic("unimplemented")
}

// StoreSession implements application.SessionsRepository
func (*SessionRepository) StoreSession(ctx context.Context, session core.Session) error {
	panic("unimplemented")
}

// DeleteSession implements application.SessionsRepository
func (*SessionRepository) DeleteSession(ctx context.Context, sessionID string) error {
	panic("unimplemented")
}
