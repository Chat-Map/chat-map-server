package application

import (
	"github.com/lordvidex/errs"
)

var (
	errNotChatMember = func(chatID, userID int64) error {
		return errs.B().Code(errs.Forbidden).Msgf("user %d is not a member of chat %d", userID, chatID).Err()
	}
)
