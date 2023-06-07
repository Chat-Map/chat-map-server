package api

import (
	"encoding/json"
	"net/http"

	"github.com/Chat-Map/chat-map-server/internal/application"
	"github.com/gorilla/mux"
)

func (s *Server) search(w http.ResponseWriter, r *http.Request) {
	// Get pattern from var
	vars := mux.Vars(r)
	pattern := vars["pattern"]
	// Do request
	usersBySearch, err := s.uc.SearchUserByAll.Execute(r.Context(), application.SearchCommandRequest{Pattern: pattern})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	// Write response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usersBySearch)
}
