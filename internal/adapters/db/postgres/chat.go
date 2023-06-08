package postgres

import (
	"context"
	"database/sql"

	"github.com/Chat-Map/chat-map-server/internal/adapters/db/postgres/sqlc"
	"github.com/Chat-Map/chat-map-server/internal/core"
	"github.com/lordvidex/errs"
)

type ChatRepository struct {
	q  *sqlc.Queries
	db *sql.DB
}

func NewChatRepository(db *sql.DB) *ChatRepository {
	return &ChatRepository{db: db, q: sqlc.New()}
}

// CreatePrivateChat implements application.ChatRepository
func (cr *ChatRepository) CreatePrivateChat(ctx context.Context, userIDs []int64) (int64, error) {
	// Begin tx
	tx, err := cr.db.Begin()
	if err != nil {
		return 0, errorTxNotStarted(err)
	}
	defer rollback(tx)
	// DO
	id, err := cr.q.CreateChat(ctx, tx, sqlc.ChatTPrivate)
	if err != nil {
		return 0, errs.B(err).Code(errs.NotFound).Msg("failed to create chat").Err()
	}
	for _, uID := range userIDs {
		err = cr.q.AddChatMember(ctx, tx, sqlc.AddChatMemberParams{ChatID: id, UserID: uID})
		if err != nil {
			return 0, errs.B(err).Code(errs.Internal).Msg("failed to add chat member").Err()
		}
	}
	// Commit
	err = tx.Commit()
	if err != nil {
		return 0, errorTxCommitted(err)
	}
	return id, nil
}

// GetChat implements application.ChatRepository
func (cr *ChatRepository) GetChat(ctx context.Context, chatID int64) (core.Chat, error) {
	// Begin tx
	tx, err := cr.db.Begin()
	if err != nil {
		return core.Chat{}, errorTxNotStarted(err)
	}
	defer rollback(tx)
	// DO
	messages, err := cr.q.GetChatMessages(ctx, tx, chatID)
	if err != nil {
		return core.Chat{}, errs.B(err).Code(errs.Internal).Msg("failed to get chat messages").Err()
	}
	users, err := cr.q.GetChatMembers(ctx, tx, chatID)
	if err != nil {
		return core.Chat{}, errs.B(err).Code(errs.Internal).Msg("failed to get chat members").Err()
	}
	// Commit
	err = tx.Commit()
	if err != nil {
		return core.Chat{}, errorTxCommitted(err)
	}
	// Return
	msgs := make([]core.Message, len(messages))
	for i, m := range messages {
		msgs[i] = convertMessage(m)
	}
	return core.Chat{
		ID:       chatID,
		UserIDs:  users,
		Messages: msgs,
	}, nil
}

// GetChatMetadata implements application.ChatRepository
func (cr *ChatRepository) GetChatsMetadata(ctx context.Context, userID int64) ([]core.ChatMetaData, error) {
	// Begin tx
	tx, err := cr.db.Begin()
	if err != nil {
		return nil, errorTxNotStarted(err)
	}
	defer rollback(tx)
	// DO
	res, err := cr.q.GetUserChatMetadata(ctx, tx, userID)
	if err != nil {
		return nil, errs.B(err).Code(errs.Internal).Msg("failed to get chat metadata").Err()
	}
	// Commit
	err = tx.Commit()
	if err != nil {
		return nil, errorTxCommitted(err)
	}
	// Return
	metadata := make([]core.ChatMetaData, len(res))
	for i, m := range res {
		metadata[i] = convertMetadata(m)
	}
	return metadata, nil
}

func (cr *ChatRepository) IsChatMember(ctx context.Context, chatID int64, userID int64) (bool, error) {
	// Begin tx
	tx, err := cr.db.Begin()
	if err != nil {
		return false, errorTxNotStarted(err)
	}
	defer rollback(tx)
	// DO
	row, err := cr.q.GetChatUserRow(ctx, tx, sqlc.GetChatUserRowParams{
		ChatID: chatID,
		UserID: userID,
	})
	if err != nil {
		return false, errs.B(err).Code(errs.Internal).Msg("failed to get chat member row").Err()
	}
	if row.ChatID == 0 || row.UserID == 0 {
		return false, errs.B().Code(errs.NotFound).Msg("chat member not found").Err()
	}
	// Commit
	err = tx.Commit()
	if err != nil {
		return false, errorTxCommitted(err)
	}
	return true, nil
}

func convertMessage(m sqlc.Message) core.Message {
	return core.Message{
		ID:        m.ID,
		Content:   m.Content,
		SenderID:  m.SenderID,
		CreatedAt: m.CreatedAt,
	}
}

func convertMetadata(md sqlc.GetUserChatMetadataRow) core.ChatMetaData {
	return core.ChatMetaData{
		ID:            md.ID,
		FirstName:     md.FirstName,
		LastName:      md.LastName,
		LatestMessage: md.LastMessage,
	}
}
