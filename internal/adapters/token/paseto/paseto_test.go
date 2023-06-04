package paseto

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/Chat-Map/chat-map-server/internal/core"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewPaseto(t *testing.T) {
	test := []struct {
		name    string
		key     string
		wantErr bool
	}{
		{
			name:    "sucess",
			key:     "12345678901234567890123456789012",
			wantErr: false,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewPaseto([]byte(tt.key))
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestPasetoTokenizerGenerateToken(t *testing.T) {
	ctx := context.Background()
	// Create paseto
	p, err := NewPaseto([]byte("12345678901234567890123456789012"))
	require.NoError(t, err)
	require.NotNil(t, p)

	tests := []struct {
		name        string
		payload     core.Payload
		validateErr bool
	}{
		{
			name: "success",
			payload: core.Payload{
				UserID:    123,
				SessionID: uuid.New(),
				ExpiresAt: time.Now().Add(time.Hour),
				CreatedAt: time.Now(),
			},
			validateErr: false,
		},
		{
			name: "expired",
			payload: core.Payload{
				UserID:    123,
				SessionID: uuid.New(),
				ExpiresAt: time.Now().Add(-time.Hour),
				CreatedAt: time.Now(),
			},
			validateErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Generate token
			token, err := p.GenerateToken(ctx, test.payload)
			require.NoError(t, err)
			require.NotEmpty(t, t)

			// Validate token
			payload, err := p.ValidateToken(ctx, token)
			if test.validateErr {
				require.Error(t, err)
				return
			}

			// Compare payload
			require.NoError(t, err)
			require.Equal(t, test.payload.UserID, payload.UserID)
			require.WithinDuration(t, test.payload.ExpiresAt, payload.ExpiresAt, time.Second)
			require.WithinDuration(t, test.payload.CreatedAt, payload.CreatedAt, time.Second)
			require.True(t, reflect.DeepEqual(test.payload.SessionID, payload.SessionID))
		})
	}
}
