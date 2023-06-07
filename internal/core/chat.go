package core

import "time"

type Chat struct {
	ID       int64
	UserIDs  []int64
	Messages []Message
}

type ChatMetaData struct {
	ID            int64
	FirstName     string
	LastName      string
	LatestMessage string
}

type Message struct {
	ID        int64
	Content   string
	SenderID  int64
	CreatedAt time.Time
}

type NotifyChat struct {
	ChatID int64
}
