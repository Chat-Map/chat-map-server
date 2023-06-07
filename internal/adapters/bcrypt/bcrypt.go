package bcrypt

import (
	"context"

	"github.com/lordvidex/errs"
	"golang.org/x/crypto/bcrypt"
)

const minPasswordLen = 8

type BcryptHasher struct{}

func New() *BcryptHasher {
	return &BcryptHasher{}
}

func (c *BcryptHasher) Hash(_ context.Context, password string) (string, error) {
	if password == "" {
		return "", errs.B().Code(errs.InvalidArgument).Msgf("password length is less than: %d", minPasswordLen).Err()
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errs.B(err).Code(errs.Unauthenticated).Msg("failed to hash password:").Err()
	}
	return string(hash), nil
}

func (c *BcryptHasher) Compare(_ context.Context, hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
