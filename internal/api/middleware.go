package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/Chat-Map/chat-map-server/internal/application"
	"github.com/Chat-Map/chat-map-server/internal/core"
	"github.com/lordvidex/errs"
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
			newFailureResponse("authorization header not provided", errs.B().Code(errs.Unauthenticated).Err()).Write(w)
			return
		}
		// Check authorization type
		if strings.HasPrefix(value, authorizationType+" ") {
			newFailureResponse("authorization type not provided", errs.B().Code(errs.Unauthenticated).Err()).Write(w)
			return
		}
		// Validate token
		token := value[len(authorizationType):]
		payload, err := s.uc.ValidateToken.Execute(r.Context(), application.TokenValidateCommandRequest{Token: token})
		if err != nil {
			newFailureResponse("failed to execute", err).Write(w)
			return
		}
		// Set payload into context as variable
		ctx := context.WithValue(r.Context(), core.PayloadKey, payload)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
