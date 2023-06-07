package application

import (
	"context"
	"testing"

	"github.com/Chat-Map/chat-map-server/internal/mock"
	"github.com/golang/mock/gomock"
)

func TestSearchCommandImplV1Execute(t *testing.T) {
	type args struct {
		params SearchCommandRequest
	}
	tests := []struct {
		name  string
		args  args
		check func(t *testing.T, response SearchCommandResponse, err error)
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			v := mock.NewMockValidator(ctrl)
			ur := mock.NewMockUserRepository(ctrl)

			s := NewSearchCommandImplV1(v, ur)

			got, err := s.Execute(context.TODO(), tt.args.params)
			tt.check(t, got, err)
		})
	}
}
