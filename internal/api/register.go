package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Chat-Map/chat-map-server/internal/application"
)

func (s *Server) register(w http.ResponseWriter, r *http.Request) {
	var body application.SignupCommandRequest
	// Read body
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
		return
	}
	// Do request
	err = s.uc.Signup.Execute(s.ctx, body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	// Write response
	w.WriteHeader(http.StatusOK)
}
