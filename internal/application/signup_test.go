package application

import (
	"context"
	"testing"

	"github.com/Chat-Map/chat-map-server/internal/mock"
	"github.com/golang/mock/gomock"
)

func TestSignupCommandImplV1Execute(t *testing.T) {
	type args struct {
		params SignupCommandRequest
	}

	tests := []struct {
		name  string
		args  args
		check func(t *testing.T, err error)
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			v := mock.NewMockValidator(ctrl)
			ur := mock.NewMockUserRepository(ctrl)
			ph := mock.NewMockPasswordHasher(ctrl)

			s := NewSignupCommandImplV1(v, ur, ph)

			err := s.Execute(context.TODO(), tt.args.params)
			tt.check(t, err)
		})
	}
}
