package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Chat-Map/chat-map-server/internal/application"
	"github.com/gorilla/mux"
)

func (s *Server) chatGet(w http.ResponseWriter, r *http.Request) {
	// Get chatID from var
	vars := mux.Vars(r)
	id := vars["id"]
	chatID, err := strconv.Atoi(id)
	if err != nil || chatID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid chat id or not provided"))
		return
	}
	// Do request
	chat, err := s.uc.ChatGet.Execute(r.Context(), application.GetChatCommandRequest{ChatID: int64(chatID)})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	// Write response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chat)
}
