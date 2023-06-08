package postgres

import (
	"context"
	"database/sql"

	"github.com/Chat-Map/chat-map-server/internal/adapters/db/postgres/sqlc"
	"github.com/Chat-Map/chat-map-server/internal/core"
	"github.com/google/uuid"
	"github.com/lordvidex/errs"
)

type SessionRepository struct {
	q  *sqlc.Queries
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{db: db, q: sqlc.New()}
}

// GetSession implements application.SessionsRepository
func (sr *SessionRepository) GetSession(ctx context.Context, sessionID uuid.UUID) (core.Session, error) {
	// Begin tx
	tx, err := sr.db.Begin()
	if err != nil {
		return core.Session{}, errorTxNotStarted(err)
	}
	defer rollback(tx)
	// Do
	res, err := sr.q.GetSession(ctx, tx, sessionID)
	if err != nil {
		return core.Session{}, errs.B(err).Code(errs.NotFound).Msg("failed to get session").Err()
	}
	// Commit
	err = tx.Commit()
	if err != nil {
		return core.Session{}, errorTxCommitted(err)
	}
	// Return
	session := convertSession(res)
	return session, nil
}

// StoreSession implements application.SessionsRepository
func (sr *SessionRepository) StoreSession(ctx context.Context, session core.Session) error {
	// Begin tx
	tx, err := sr.db.Begin()
	if err != nil {
		return errorTxNotStarted(err)
	}
	defer rollback(tx)
	// Do
	err = sr.q.StoreSession(ctx, tx, sqlc.StoreSessionParams{
		ID:        session.ID,
		UserID:    session.UserID,
		ExpiresAt: session.ExpiresAt,
	})
	if err != nil {
		return errs.B(err).Code(errs.Internal).Msg("failed to store session").Err()
	}
	// Commit
	err = tx.Commit()
	if err != nil {
		return errorTxCommitted(err)
	}
	// Return
	return nil
}

func convertSession(s sqlc.Session) core.Session {
	return core.Session{
		ID:        s.ID,
		UserID:    s.UserID,
		ExpiresAt: s.ExpiresAt,
	}
}
