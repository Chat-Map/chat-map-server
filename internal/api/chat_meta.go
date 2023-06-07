package api

import (
	"encoding/json"
	"net/http"

	"github.com/Chat-Map/chat-map-server/internal/application"
)

func (s *Server) chatGetMeta(w http.ResponseWriter, r *http.Request) {
	// Do request
	metadata, err := s.uc.ChatMeta.Execute(r.Context(), application.GetChatMetaCommandRequest{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	// Write response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(metadata)
}
