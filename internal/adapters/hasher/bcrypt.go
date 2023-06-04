package hasher

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const minPasswordLen = 8

type BcryptHasher struct{}

func NewBcryptHasher() *BcryptHasher {
	return &BcryptHasher{}
}

func (c *BcryptHasher) Hash(_ context.Context, password string) (string, error) {
	if password == "" {
		return "", fmt.Errorf("password length is less than %d", minPasswordLen)
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hash), nil
}

func (c *BcryptHasher) Compare(_ context.Context, hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
