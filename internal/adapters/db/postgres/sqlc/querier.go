// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddChatMember(ctx context.Context, db DBTX, arg AddChatMemberParams) error
	CreateChat(ctx context.Context, db DBTX, chatType ChatT) (int64, error)
	GetAllUsers(ctx context.Context, db DBTX) ([]GetAllUsersRow, error)
	GetChatMembers(ctx context.Context, db DBTX, chatID int64) ([]int64, error)
	GetChatMessages(ctx context.Context, db DBTX, chatID int64) ([]Message, error)
	GetChatUserRow(ctx context.Context, db DBTX, arg GetChatUserRowParams) (ChatUser, error)
	GetSession(ctx context.Context, db DBTX, id uuid.UUID) (Session, error)
	GetUserByEmail(ctx context.Context, db DBTX, email string) (User, error)
	GetUserByID(ctx context.Context, db DBTX, id int64) (User, error)
	GetUserChatMetadata(ctx context.Context, db DBTX, id int64) ([]GetUserChatMetadataRow, error)
	SearchUserByAll(ctx context.Context, db DBTX, pattern string) ([]SearchUserByAllRow, error)
	StoreMessage(ctx context.Context, db DBTX, arg StoreMessageParams) (int64, error)
	StoreSession(ctx context.Context, db DBTX, arg StoreSessionParams) error
	StoreUser(ctx context.Context, db DBTX, arg StoreUserParams) error
}

var _ Querier = (*Queries)(nil)
