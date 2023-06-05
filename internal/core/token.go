package core

import (
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	UserID    int64     `json:"user_id"`
	SessionID uuid.UUID `json:"session_id"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}
