package core

import "time"

type Chat struct {
	UserID   string    `json:"user_id"`
	Messages []Message `json:"messages"`
}

type ChatMetaData struct {
	ID            string  `json:"id"`
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	LatestMessage Message `json:"latest_message"`
}

type Message struct {
	ID        int64     `json:"id,omitempty"`
	Text      string    `json:"text"`
	FromUser  string    `json:"from"` // User ID
	CreatedAt time.Time `json:"created_at,omitempty"`
}
