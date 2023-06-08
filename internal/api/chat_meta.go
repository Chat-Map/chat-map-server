package api

import (
	"net/http"

	"github.com/Chat-Map/chat-map-server/internal/application"
)

type chatMetaGetResponseDTO struct {
	ID            int64  `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	LatestMessage string `json:"latest_message"`
}

func (chatMetaGetResponseDTO) from(x application.GetChatMetaCommandResponse) []chatMetaGetResponseDTO {
	res := make([]chatMetaGetResponseDTO, len(x.ChatMetaData))
	for i, v := range x.ChatMetaData {
		res[i] = chatMetaGetResponseDTO{
			ID:            v.ID,
			FirstName:     v.FirstName,
			LastName:      v.LastName,
			LatestMessage: v.LatestMessage,
		}
	}
	return res
}

// ChatMetadatGet godoc
//
//	@Summary		Get user's chats metadata
//	@Description	Get user's chats metadata for a given user(User ID is taken from the token payload)
//	@Tags			chat
//	@Accept			json
//	@Produce		json
//	@Success		200					{object}	api.Response{data=[]chatMetaGetResponseDTO}
//	@Failure		400,401,403,404,500	{object}	api.Response{data=interface{}}
//	@Security		bearerAuth
//	@Router			/chat/meta [get]
func (s *Server) chatGetMeta(w http.ResponseWriter, r *http.Request) {
	// Do request
	metadata, err := s.uc.ChatMeta.Execute(r.Context(), application.GetChatMetaCommandRequest{})
	if err != nil {
		newFailureResponse("failed to execute", err).Write(w)
		return
	}
	// Write response
	newSuccessResponse("metadata fetched", new(chatMetaGetResponseDTO).from(metadata)).Write(w)
}
