package core

import (
	"github.com/lordvidex/errs"
)

type ContextKey string

var (
	PayloadKey ContextKey = "payload"
)

var (
	ErrNoPayload = errs.B().Code(errs.Unauthenticated).Msg("no payload is set").Err()
)
