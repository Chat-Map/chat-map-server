package notifier

import (
	"context"
	"fmt"

	"github.com/Chat-Map/chat-map-server/internal/core"
)

type Notifier struct {
	userToAddr map[int64]map[string]bool
	addrToChan map[string]chan core.NotifyChat
}

func New() *Notifier {
	return &Notifier{
		userToAddr: make(map[int64]map[string]bool),
		addrToChan: make(map[string]chan core.NotifyChat),
	}
}

// Listen listens for notifications on the provided address and returns
// the corresponding channel. Returns an error if no channel is found for
// the specified address. i.e register before calling this function
//
// address: The address to listen for notifications.
// core.NotifyChat is returned, along with an error if there is one.
func (n *Notifier) Listen(ctx context.Context, address string) (core.NotifyChat, error) {
	if _, ok := n.addrToChan[address]; !ok {
		return core.NotifyChat{}, fmt.Errorf("no channel found for address: %s", address)
	}
	return <-n.addrToChan[address], nil
}

// Notify sends a notification to all users in userIDs that a new chat has been created
// the given chat id
//
// userIDs: A slice of int64 user IDs to notify.
// chatID: The int64 ID of the newly created chat to notify users about.
func (n *Notifier) Notify(ctx context.Context, userIDs []int64, chatID int64) {
	for _, userID := range userIDs {
		for ip := range n.userToAddr[userID] {
			go func(ip string) {
				n.addrToChan[ip] <- core.NotifyChat{ChatID: chatID}
			}(ip)
		}
	}
}

// Register adds the given address to the list of addresses that notifications
// will be sent to when userID is notified. If the address wasn't already
// registered, a new channel is created for it.
//
// userID: The ID of the user to register the address for.
// address: The address to register.
func (n *Notifier) Register(ctx context.Context, userID int64, address string) {
	if _, ok := n.userToAddr[userID]; !ok {
		n.userToAddr[userID] = make(map[string]bool)
	}
	n.userToAddr[userID][address] = true
	if _, ok := n.addrToChan[address]; !ok {
		n.addrToChan[address] = make(chan core.NotifyChat)
	}
}

// Unregister removes the provided address from the list of addresses
// associated with the provided userID in the Notifier.
//
// userID: int64 value representing the user ID.
// address: string value representing the address to be deleted.
func (n *Notifier) Unregister(ctx context.Context, userID int64, address string) {
	if _, ok := n.userToAddr[userID]; !ok {
		return
	}
	close(n.addrToChan[address])
	delete(n.addrToChan, address)
	delete(n.userToAddr[userID], address)
}
