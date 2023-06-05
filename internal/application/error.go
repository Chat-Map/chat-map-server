package application

import "fmt"

var (
	errNotChatMember = func(chatID, userID int64) error {
		return fmt.Errorf("user %d is not a member of chat %d", userID, chatID)
	}
)
