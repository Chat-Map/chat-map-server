package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Chat-Map/chat-map-server/internal/application"
	"github.com/Chat-Map/chat-map-server/internal/core"
)

type loginRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponseDTO struct {
	User         core.User `json:"user"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func (loginResponseDTO) from(x application.SigninCommandResponse) loginResponseDTO {
	return loginResponseDTO{
		User:         x.User,
		AccessToken:  x.AccessToken,
		RefreshToken: x.RefreshToken,
		ExpiresAt:    x.ExpiresAt,
	}
}

func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	var body loginRequestDTO
	// Read body
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		newFailureResponse("failed to read body", err).Write(w)
		return
	}
	// Close body
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("failed to close body on request: %+v", err)
		}
	}()
	// Unmarshal body
	err = json.Unmarshal(bytes, &body)
	if err != nil {
		newFailureResponse("failed to unmarshal body", err).Write(w)
		return
	}
	// Do request
	response, err := s.uc.Signin.Execute(s.ctx, application.SigninCommandRequest{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		newFailureResponse("failed to execute", err).Write(w)
		return
	}
	// Write response
	newSuccessResponse("logged in", new(loginResponseDTO).from(response)).Write(w)
}
