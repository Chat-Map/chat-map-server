package application

import (
	"context"
	"testing"

	"github.com/Chat-Map/chat-map-server/internal/mock"
	"github.com/golang/mock/gomock"
)

func TestCreateChatCommandImplV1Execute(t *testing.T) {
	type args struct {
		params CreateChatCommandRequest
	}

	tests := []struct {
		name  string
		args  args
		check func(t *testing.T, id int32, err error)
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			v := mock.NewMockValidator(ctrl)
			cr := mock.NewMockChatRepository(ctrl)

			s := NewCreateChatCommandImplV1(v, cr)

			id, err := s.Execute(context.TODO(), tt.args.params)
			tt.check(t, id, err)
		})
	}
}
