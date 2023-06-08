package application

import (
	"context"
	"testing"

	"github.com/Chat-Map/chat-map-server/internal/mock"
	"github.com/golang/mock/gomock"
)

func TestGetChatMetaCommandImplV1Execute(t *testing.T) {
	type args struct {
		params GetChatMetaCommandRequest
	}

	tests := []struct {
		name  string
		args  args
		check func(t *testing.T, metadata GetChatMetaCommandResponse, err error)
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			cr := mock.NewMockChatRepository(ctrl)

			s := NewGetChatMetaCommandImplV1(cr)

			metadata, err := s.Execute(context.TODO(), tt.args.params)
			tt.check(t, metadata, err)
		})
	}
}
