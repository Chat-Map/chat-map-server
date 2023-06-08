package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Chat-Map/chat-map-server/internal/application"
)

// Register godoc
//
//	@Summary		Notify the client when a new chat is created with him
//	@Description	Notify the client when a new chat is created with him and send to connect to the room
//	@Tags			chat
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	api.Response{data=int64}
//	@Router			/chat/notify [get]
func (s *Server) notify(w http.ResponseWriter, r *http.Request) {
	// Set the response headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	listner, closeFn := s.uc.NotifyListen.Execute(r.Context(), application.NotifyListenRequest{Address: r.RemoteAddr})
	defer closeFn()
	for message := range listner {
		// Send notification about the newly created chat
		// message := nc.receive(r.RemoteAddr)
		// TODO: get message from the passed notifier
		err := json.NewEncoder(w).Encode(message)
		if err != nil {
			log.Printf("error sending notification to %s: %v", r.RemoteAddr, err)
			return
		}
		flusher := w.(http.Flusher)
		flusher.Flush()
	}
}
