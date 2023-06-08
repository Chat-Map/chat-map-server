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

type registerRequestDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type registerResponseDTO struct {
	User         core.User `json:"user"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func (registerResponseDTO) from(x application.SigninCommandResponse) registerResponseDTO {
	return registerResponseDTO{
		User:         x.User,
		AccessToken:  x.AccessToken,
		RefreshToken: x.RefreshToken,
		ExpiresAt:    x.ExpiresAt,
	}
}

// Register godoc
//
//	@Summary		Register a new user account
//	@Description	Register a new user account with the given information
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body				body		registerRequestDTO	true	"User ID"
//	@Success		200					{object}	api.Response{data=registerResponseDTO}
//	@Failure		400,401,403,404,500	{object}	api.Response{data=interface{}}
//	@Router			/auth/register [post]
func (s *Server) register(w http.ResponseWriter, r *http.Request) {
	var body registerRequestDTO
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
	err = s.uc.Signup.Execute(s.ctx, application.SignupCommandRequest{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Phone:     body.Phone,
		Email:     body.Email,
		Password:  body.Password,
	})
	if err != nil {
		newFailureResponse("failed to execute signup", err).Write(w)
		return
	}
	res, err := s.uc.Signin.Execute(s.ctx, application.SigninCommandRequest{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		newFailureResponse("failed to execute signin", err).Write(w)
		return
	}
	// Write response
	newSuccessResponse("registered", new(registerResponseDTO).from(res)).Write(w)
}
