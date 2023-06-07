package core

import "time"

type Chat struct {
	ID       int64     `json:"id"`
	UserIDs  []int64   `json:"user_id"`
	Messages []Message `json:"messages"`
}

type ChatMetaData struct {
	ID            int64  `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	LatestMessage string `json:"latest_message"`
}

type Message struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	SenderID  int64     `json:"sender_id"`
	CreatedAt time.Time `json:"created_at"`
}

type NotifyChat struct {
	ChatID int64 `json:"chat_id"`
}
