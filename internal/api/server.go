package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Chat-Map/chat-map-server/internal/application"
	"github.com/gorilla/mux"
)

type Server struct {
	ctx context.Context
	uc  *application.UseCase
	r   *mux.Router
}

func NewServer(ctx context.Context, uc *application.UseCase) *Server {
	s := &Server{ctx: ctx, uc: uc}
	s.setup()
	return s
}

func (s *Server) setup() {
	r := mux.NewRouter()

	rr := r.PathPrefix("/api/v1").Subrouter()

	// Register
	rr.HandleFunc("/register", s.register).Methods("POST")
	rr.HandleFunc("/login", s.login).Methods("POST")

	// Search
	rr.HandleFunc("/search/{pattern}", s.authMW(s.search)).Methods("GET") // Search for User/Group/Channel

	// Chat
	rr.HandleFunc("/chat", s.authMW(s.chatCreate)).Methods("POST")
	rr.HandleFunc("/chat/{id}", s.authMW(s.chatGet)).Methods("GET")
	rr.HandleFunc("/chat/ws/{id}", s.authMW(s.chatws)).Methods("GET")
	rr.HandleFunc("/chat/meta", s.authMW(s.chatGetMeta)).Methods("GET")

	s.r = r
}

func (s *Server) Run(port string) error {
	log.Printf("Starting server at port: %s", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), s.r)
}
