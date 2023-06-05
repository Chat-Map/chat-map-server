package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/Chat-Map/chat-map-server/internal/application"
	"github.com/Chat-Map/chat-map-server/internal/core"
)

const (
	authorizationHeader = "authorization"
	authorizationType   = "Bearer"
)

func (s *Server) authMW(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		value := r.Header.Get(authorizationHeader)
		// Check authorization header
		if value == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("authorization header not provided"))
			return
		}
		// Check authorization type
		if strings.HasPrefix(value, authorizationType+" ") {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("authorization type not provided"))
			return
		}
		// Validate token
		token := value[len(authorizationType):]
		payload, err := s.uc.ValidateToken.Execute(r.Context(), application.TokenValidateCommandRequest{Token: token})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		// Set payload into context as variable
		ctx := context.WithValue(r.Context(), core.PayloadKey, payload)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
