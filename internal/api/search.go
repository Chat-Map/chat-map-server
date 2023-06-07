package api

import (
	"net/http"

	"github.com/Chat-Map/chat-map-server/internal/application"
	"github.com/gorilla/mux"
)

type searchResponseDTO struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (searchResponseDTO) from(x application.SearchCommandResponse) []searchResponseDTO {
	res := make([]searchResponseDTO, len(x.Users))
	for i, v := range x.Users {
		res[i] = searchResponseDTO{
			ID:        v.ID,
			FirstName: v.FirstName,
			LastName:  v.LastName,
		}
	}
	return res
}

func (s *Server) search(w http.ResponseWriter, r *http.Request) {
	// Get pattern from var
	vars := mux.Vars(r)
	pattern := vars["pattern"]
	// Do request
	usersBySearch, err := s.uc.SearchUserByAll.Execute(r.Context(), application.SearchCommandRequest{Pattern: pattern})
	if err != nil {
		newFailureResponse("failed to execute", err).Write(w)
		return
	}
	// Write response
	newSuccessResponse("found users by search", new(searchResponseDTO).from(usersBySearch)).Write(w)
}
