package application

import (
	"context"
	"testing"

	"github.com/Chat-Map/chat-map-server/internal/mock"
	"github.com/golang/mock/gomock"
)

func TestSigninCommandImplV1Execute(t *testing.T) {
	type args struct {
		params SigninCommandRequest
	}

	tests := []struct {
		name  string
		args  args
		want  SigninCommandResponse
		check func(res SigninCommandResponse, err error)
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			v := mock.NewMockValidator(ctrl)
			ur := mock.NewMockUserRepository(ctrl)
			sr := mock.NewMockSessionsRepository(ctrl)
			ph := mock.NewMockPasswordHasher(ctrl)
			tk := mock.NewMockTokenizer(ctrl)

			s := NewSigninCommandImplV1(v, ur, sr, ph, tk)

			got, err := s.Execute(context.TODO(), tt.args.params)
			tt.check(got, err)
		})
	}
}
