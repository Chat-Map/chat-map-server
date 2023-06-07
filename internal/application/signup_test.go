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
		check func(t *testing.T, res SignupCommandResponse, err error)
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
			sr := mock.NewMockSessionsRepository(ctrl)
			tk := mock.NewMockTokenizer(ctrl)

			signin := NewSigninCommandImplV1(v, ur, sr, ph, tk)
			signup := NewSignupCommandImplV1(v, ur, ph, signin)

			res, err := signup.Execute(context.TODO(), tt.args.params)
			tt.check(t, res, err)
		})
	}
}
