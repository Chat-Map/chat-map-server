package postgres

import (
	"context"
	"database/sql"

	"github.com/Chat-Map/chat-map-server/internal/core"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetUser implements application.UserRepository
func (*UserRepository) GetUser(ctx context.Context, userID string) (core.User, error) {
	panic("unimplemented")
}

// SearchUser implements application.UserRepository
func (*UserRepository) SearchUserByEmail(ctx context.Context, email string) ([]core.UserBySearch, error) {
	panic("unimplemented")
}

// StoreUser implements application.UserRepository
func (*UserRepository) StoreUser(ctx context.Context, user core.User) error {
	panic("unimplemented")
}
