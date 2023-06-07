package paseto

import (
	"context"
	"time"

	"github.com/Chat-Map/chat-map-server/internal/core"
	"github.com/lordvidex/errs"
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
		return nil, errs.B().Code(errs.InvalidArgument).Msgf("secretKey len is less than the min value %d", minSecretKeyLen).Err()
	}
	// Return a new paseto tokenizer
	return &PasetoTokenizer{p: paseto.NewV2(), sk: secretKey}, nil
}

// generateToken Create a new token with user, sessionID, exp(Token life duration)
func (pt *PasetoTokenizer) GenerateToken(ctx context.Context, payload core.Payload) (string, error) {
	token, err := pt.p.Encrypt(pt.sk, payload, nil)
	if err != nil {
		return "", errs.B(err).Code(errs.Internal).Msg("failed to create token").Err()
	}
	return token, nil
}

// DecryptToken decrypts the token to get `TokenPayload` & verifies that token hasn't expired
func (pt *PasetoTokenizer) ValidateToken(ctx context.Context, token string) (core.Payload, error) {
	var payload core.Payload
	err := pt.p.Decrypt(token, pt.sk, &payload, nil)
	if err != nil {
		return core.Payload{}, errs.B(err).Code(errs.Internal).Msg("failed to decrypt token").Err()
	}
	if payload.ExpiresAt.Before(time.Now()) {
		return core.Payload{}, errs.B().Code(errs.Unauthenticated).Msg("token has expired").Err()
	}
	return payload, nil
}
