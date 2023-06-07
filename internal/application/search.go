package application

import (
	"context"

	"github.com/Chat-Map/chat-map-server/internal/core"
)

type SearchCommandRequest struct {
	Pattern string `validate:"required"`
}

type SearchCommandResponse struct {
	Users []core.UserBySearch
}

type SearchCommand interface {
	Execute(ctx context.Context, params SearchCommandRequest) (SearchCommandResponse, error)
}

type SearchCommandImplV1 struct {
	v  Validator
	ur UserRepository
}

func NewSearchCommandImplV1(v Validator, ur UserRepository) *SearchCommandImplV1 {
	return &SearchCommandImplV1{v: v, ur: ur}
}

func (s SearchCommandImplV1) Execute(ctx context.Context, params SearchCommandRequest) (SearchCommandResponse, error) {
	err := s.v.Validate(ctx, params)
	if err != nil {
		return SearchCommandResponse{}, err
	}
	user, err := s.ur.SearchUserByAll(ctx, params.Pattern)
	if err != nil {
		return SearchCommandResponse{}, err
	}
	return SearchCommandResponse{
		Users: user,
	}, nil
}
