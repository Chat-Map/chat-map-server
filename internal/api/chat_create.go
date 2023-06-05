package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Chat-Map/chat-map-server/internal/application"
)

func (s *Server) chatCreate(w http.ResponseWriter, r *http.Request) {
	var body application.CreateChatCommandRequest
	// Read body
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	// Do request
	id, err := s.uc.ChatCreate.Execute(r.Context(), body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	// Write response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}
