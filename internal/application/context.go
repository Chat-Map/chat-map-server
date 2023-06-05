package application

import (
	"context"

	"github.com/Chat-Map/chat-map-server/internal/core"
)

func GetPayload(ctx context.Context) (core.Payload, error) {
	payload, ok := ctx.Value(core.PayloadKey).(core.Payload)
	if !ok {
		return core.Payload{}, core.ErrNoPayload
	}
	return payload, nil
}
