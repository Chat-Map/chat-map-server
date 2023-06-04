package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Chat-Map/chat-map-server/internal/adapters/db/postgres/sqlc"
	"github.com/Chat-Map/chat-map-server/internal/core"
)

type UserRepository struct {
	q  *sqlc.Queries
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db, q: sqlc.New()}
}

// GetUser implements application.UserRepository
func (ur *UserRepository) GetUser(ctx context.Context, userID int32) (core.User, error) {
	// Begin tx
	tx, err := ur.db.Begin()
	if err != nil {
		return core.User{}, errorTxNotStarted(err)
	}
	defer rollback(tx)
	// Do
	res, err := ur.q.GetUserByID(ctx, tx, userID)
	if err != nil {
		return core.User{}, fmt.Errorf("failed to get user: %w", err)
	}
	// Commit
	err = tx.Commit()
	if err != nil {
		return core.User{}, errorTxCommitted(err)
	}
	// Return
	u := convertUser(res)
	return u, nil
}

func (ur *UserRepository) GetByEmail(ctx context.Context, email string) (core.User, error) {
	// Begin tx
	tx, err := ur.db.Begin()
	if err != nil {
		return core.User{}, errorTxNotStarted(err)
	}
	defer rollback(tx)
	// Do
	res, err := ur.q.GetUserByEmail(ctx, tx, email)
	if err != nil {
		return core.User{}, fmt.Errorf("failed to get user: %w", err)
	}
	// Commit
	err = tx.Commit()
	if err != nil {
		return core.User{}, errorTxCommitted(err)
	}
	// Return
	u := convertUser(res)
	return u, nil
}

func (ur *UserRepository) GetAllUsers(ctx context.Context) ([]core.UserBySearch, error) {
	// Begin tx
	tx, err := ur.db.Begin()
	if err != nil {
		return nil, errorTxNotStarted(err)
	}
	defer rollback(tx)
	// Do
	res, err := ur.q.GetAllUsers(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	// Commit
	err = tx.Commit()
	if err != nil {
		return nil, errorTxCommitted(err)
	}
	// Return
	users := make([]core.UserBySearch, len(res))
	for i, u := range res {
		users[i] = convertUserBySearch(u)
	}
	return users, nil
}

// SearchUser implements application.UserRepository
func (ur *UserRepository) SearchUserByEmail(ctx context.Context, email string) ([]core.UserBySearch, error) {
	// Begin tx
	tx, err := ur.db.Begin()
	if err != nil {
		return nil, errorTxNotStarted(err)
	}
	defer rollback(tx)
	// Do

	// Commit
	err = tx.Commit()
	if err != nil {
		return nil, errorTxCommitted(err)
	}

	return nil, nil
}

// StoreUser implements application.UserRepository
func (ur *UserRepository) StoreUser(ctx context.Context, user core.User) error {
	// Begin tx
	tx, err := ur.db.Begin()
	if err != nil {
		return errorTxNotStarted(err)
	}
	defer rollback(tx)
	// Do
	err = ur.q.StoreUser(ctx, tx, sqlc.StoreUserParams{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	})
	if err != nil {
		return fmt.Errorf("failed to store user: %+v", err)
	}
	// Commit
	err = tx.Commit()
	if err != nil {
		return errorTxCommitted(err)
	}

	return nil
}

func convertUser(u sqlc.User) core.User {
	return core.User{
		ID:        u.ID,
		Phone:     u.Phone,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Password:  u.Password,
		CreatedAt: u.CreatedAt.Time,
		UpdatedAt: u.UpdatedAt.Time,
	}
}

func convertUserBySearch(u sqlc.GetAllUsersRow) core.UserBySearch {
	return core.UserBySearch{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}
