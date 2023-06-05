package application

import (
	"context"
	"testing"

	"github.com/Chat-Map/chat-map-server/internal/core"
	"github.com/Chat-Map/chat-map-server/internal/mock"
	"github.com/golang/mock/gomock"
)

func TestTokenValidateCommandImplV1Execute(t *testing.T) {
	type args struct {
		params TokenValidateCommandRequest
	}
	tests := []struct {
		name  string
		args  args
		check func(t *testing.T, c core.Payload, err error)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			v := mock.NewMockValidator(ctrl)
			tk := mock.NewMockTokenizer(ctrl)

			s := NewTokenValidateCommandImplV1(v, tk)

			payload, err := s.Execute(context.TODO(), tt.args.params)
			tt.check(t, payload, err)
		})
	}
}
