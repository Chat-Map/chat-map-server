package application

import (
	"context"
	"testing"

	"github.com/Chat-Map/chat-map-server/internal/mock"
	"github.com/golang/mock/gomock"
)

func TestStoreMessageCommandImplV1Execute(t *testing.T) {
	type args struct {
		params StoreMessageCommandRequest
	}

	tests := []struct {
		name  string
		args  args
		check func(messageID int32, err error)
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			v := mock.NewMockValidator(ctrl)
			ur := mock.NewMockUserRepository(ctrl)
			sr := mock.NewMockSessionsRepository(ctrl)
			cr := mock.NewMockChatRepository(ctrl)
			mr := mock.NewMockMessageRepository(ctrl)

			s := NewStoreMessageCommandImplV1(v, ur, cr, mr, sr)

			got, err := s.Execute(context.TODO(), tt.args.params)
			tt.check(got, err)
		})
	}
}
