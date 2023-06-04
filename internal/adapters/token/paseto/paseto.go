package paseto

import (
	"context"
	"fmt"
	"time"

	"github.com/Chat-Map/chat-map-server/internal/core"
	"github.com/o1egl/paseto"
)

const minSecretKeyLen = 32

type PasetoTokenizer struct {
	p  *paseto.V2
	sk []byte
}

// NewPaseto Creates a new instance of PasetoTokenizer
func NewPaseto(secretKey []byte) (*PasetoTokenizer, error) {
	// Check that secret key is exactly equal to `minSecretKeyLen`
	if len(secretKey) < minSecretKeyLen {
		return nil, fmt.Errorf("secretKey len is less than the min value %d", minSecretKeyLen)
	}
	// Return a new paseto tokenizer
	return &PasetoTokenizer{p: paseto.NewV2(), sk: secretKey}, nil
}

// generateToken Create a new token with user, sessionID, exp(Token life duration)
func (pt *PasetoTokenizer) GenerateToken(ctx context.Context, payload core.Payload) (string, error) {
	token, err := pt.p.Encrypt(pt.sk, payload, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create token: %+v", err)
	}
	return token, nil
}

// DecryptToken decrypts the token to get `TokenPayload` & verifies that token hasn't expired
func (pt *PasetoTokenizer) ValidateToken(ctx context.Context, token string) (core.Payload, error) {
	var payload core.Payload
	err := pt.p.Decrypt(token, pt.sk, &payload, nil)
	if err != nil {
		return core.Payload{}, fmt.Errorf("failed to decrypt token: %+v", err)
	}
	if payload.ExpiresAt.Before(time.Now()) {
		return core.Payload{}, fmt.Errorf("token has expired")
	}
	return payload, nil
}
