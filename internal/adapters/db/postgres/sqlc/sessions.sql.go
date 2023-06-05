// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: sessions.sql

package sqlc

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const getSession = `-- name: GetSession :one
SELECT id, user_id, expires_at, created_at
FROM sessions
WHERE id = $1
`

func (q *Queries) GetSession(ctx context.Context, db DBTX, id uuid.UUID) (Session, error) {
	row := db.QueryRowContext(ctx, getSession, id)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const storeSession = `-- name: StoreSession :exec
INSERT INTO sessions (id, user_id, expires_at)
VALUES ($1, $2, $3)
`

type StoreSessionParams struct {
	ID        uuid.UUID `db:"id" json:"id"`
	UserID    int64     `db:"user_id" json:"user_id"`
	ExpiresAt time.Time `db:"expires_at" json:"expires_at"`
}

func (q *Queries) StoreSession(ctx context.Context, db DBTX, arg StoreSessionParams) error {
	_, err := db.ExecContext(ctx, storeSession, arg.ID, arg.UserID, arg.ExpiresAt)
	return err
}
