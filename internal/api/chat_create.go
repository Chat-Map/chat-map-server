package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Chat-Map/chat-map-server/internal/application"
)

type chatCreateRequestDTO struct {
	UserID int64 `json:"user_id"`
}

type chatCreateResponseDTO struct {
	ChatID int64 `json:"chat_id"`
}

func (chatCreateResponseDTO) from(x application.CreateChatCommandResponse) chatCreateResponseDTO {
	return chatCreateResponseDTO{ChatID: x.ChatID}
}

// ChatCreate godoc
//
//	@Summary		Create a private chat
//	@Description	Create a private chat with a given userID
//	@Tags			chat
//	@Accept			json
//	@Produce		json
//	@Param			body			body		chatCreateRequestDTO	true	"User ID"
//	@Success		200				{object}	api.Response{data=chatCreateResponseDTO}
//	@Failure		400,401,409,500	{object}	api.Response{data=interface{}}
//	@Security		bearerAuth
//	@Router			/chat [post]
func (s *Server) chatCreate(w http.ResponseWriter, r *http.Request) {
	var body chatCreateRequestDTO
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
	res, err := s.uc.ChatCreate.Execute(r.Context(), application.CreateChatCommandRequest{UserID: body.UserID})
	if err != nil {
		newFailureResponse("failed to execute", err).Write(w)
		return
	}
	// Write response
	w.WriteHeader(http.StatusOK)
	newSuccessResponse("chat created", new(chatCreateResponseDTO).from(res)).Write(w)
}
