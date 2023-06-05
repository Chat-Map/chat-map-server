package core

import "fmt"

type ContextKey string

var (
	PayloadKey ContextKey = "payload"
)

var (
	ErrNoPayload = fmt.Errorf("no payload is set")
)
