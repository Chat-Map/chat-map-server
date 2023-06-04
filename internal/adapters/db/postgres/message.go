package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Chat-Map/chat-map-server/internal/adapters/db/postgres/sqlc"
	"github.com/Chat-Map/chat-map-server/internal/core"
)

type MessageRepository struct {
	q  *sqlc.Queries
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{db: db, q: sqlc.New()}
}

// StoreMessage implements application.MessageRepository
func (mr *MessageRepository) StoreMessage(ctx context.Context, chatID int32, message core.Message) error {
	// Begin tx
	tx, err := mr.db.Begin()
	if err != nil {
		return errorTxNotStarted(err)
	}
	defer rollback(tx)
	// Do
	err = mr.q.StoreMessage(ctx, tx, sqlc.StoreMessageParams{
		ChatID:   chatID,
		SenderID: message.SenderID,
		Content:  message.Content,
	})
	if err != nil {
		return fmt.Errorf("failed to store message: %+v", err)
	}
	// Commit
	err = tx.Commit()
	if err != nil {
		return errorTxCommitted(err)
	}
	return nil
}
