package application

import (
	"context"

	"github.com/Chat-Map/chat-map-server/internal/core"
)

type NotifyListenRequest struct {
	Address string `json:"address" validate:"required,ip"`
}

type NotifyListenResponse struct {
	Message core.NotifyChat `json:"message"`
}

type NotifyListenCommand interface {
	Execute(ctx context.Context, params NotifyListenRequest) (chan NotifyListenResponse, func())
}

type NotifyListenImplV1 struct {
	v  Validator
	cn ChatNotifier
}

func NewNotifyListenImplV1(v Validator, cn ChatNotifier) *NotifyListenImplV1 {
	return &NotifyListenImplV1{v: v, cn: cn}
}

func (s NotifyListenImplV1) Execute(ctx context.Context, params NotifyListenRequest) (chan NotifyListenResponse, func()) {
	err := s.v.Validate(ctx, params)
	if err != nil {
		return nil, nil
	}
	// Get Payload
	payload, err := GetPayload(ctx)
	if err != nil {
		return nil, nil
	}
	s.cn.Register(ctx, payload.UserID, params.Address)
	listner := s.cn.Listen(ctx, params.Address)
	// Forward listner to response channel
	response := make(chan NotifyListenResponse)
	go func() {
		for {
			message, ok := <-listner
			if !ok {
				break
			}
			response <- NotifyListenResponse{message}
		}
	}()
	// Return
	return response, func() {
		s.cn.Unregister(ctx, payload.UserID, params.Address)
		close(response)
	}
}
