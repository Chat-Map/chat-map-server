package core

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        int64
	Email     string
	Phone     string
	FirstName string
	LastName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserBySearch struct {
	ID        int64
	FirstName string
	LastName  string
}

type Session struct {
	ID        uuid.UUID
	UserID    int64
	ExpiresAt time.Time
}
