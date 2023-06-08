package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Chat-Map/chat-map-server/internal/application"
	"github.com/gorilla/mux"
)

type chatGetResponseDTO struct {
	ID       int64            `json:"id"`
	UserIDs  []int64          `json:"user_ids"`
	Messages []chatGetMessage `json:"messages"`
}

type chatGetMessage struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	SenderID  int64     `json:"sender_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (chatGetResponseDTO) from(x application.GetChatCommandResponse) chatGetResponseDTO {
	messages := make([]chatGetMessage, len(x.Chat.Messages))
	for i, v := range x.Chat.Messages {
		messages[i] = chatGetMessage{
			ID:        v.ID,
			Content:   v.Content,
			SenderID:  v.SenderID,
			CreatedAt: v.CreatedAt,
		}
	}
	return chatGetResponseDTO{
		ID:       x.Chat.ID,
		UserIDs:  x.Chat.UserIDs,
		Messages: messages,
	}
}

// ChatGet godoc
//
// @Summary		Get a specific chat all messages
// @Description	Get chat messages for a given chat ID
// @Tags			chat
// @Accept			json
// @Produce		json
// @Param			params				path		int64	true	"Chat ID"
// @Success		200					{object}	api.Response{data=chatGetResponseDTO}
// @Failure		400,401,403,404,500	{object}	api.Response{data=interface{}}
// @Security		bearerAuth
// @Router			/chat/{id} [get]
func (s *Server) chatGet(w http.ResponseWriter, r *http.Request) {
	// Get chatID from var
	vars := mux.Vars(r)
	id := vars["id"]
	chatID, err := strconv.Atoi(id)
	if err != nil || chatID == 0 {
		newFailureResponse("failed to parse chat id parameter", err).Write(w)
		return
	}
	// Do request
	chat, err := s.uc.ChatGet.Execute(r.Context(), application.GetChatCommandRequest{ChatID: int64(chatID)})
	if err != nil {
		newFailureResponse("failed to execute", err).Write(w)
		return
	}
	// Write response
	newSuccessResponse("chat fetched", new(chatGetResponseDTO).from(chat)).Write(w)
}
