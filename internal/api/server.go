package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Chat-Map/chat-map-server/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"

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
	s.swaggerDocs()
	return s
}

func (s *Server) setup() {
	r := mux.NewRouter()
	r = r.PathPrefix("/api/v1").Subrouter()

	// Register
	r.HandleFunc("/auth/register", s.register).Methods("POST")
	r.HandleFunc("/auth/login", s.login).Methods("POST")

	// User
	r.HandleFunc("/user/search/{pattern}", s.authMW(s.search)).Methods("GET") // Search for User/Group/Channel

	// Chat
	r.HandleFunc("/chat", s.authMW(s.chatCreate)).Methods("POST")
	r.HandleFunc("/chat/meta", s.authMW(s.chatGetMeta)).Methods("GET")
	r.HandleFunc("/chat/{id}", s.authMW(s.chatGet)).Methods("GET")

	// Chat (websockets & sse)
	r.HandleFunc("/chat/ws/{id}", s.authMW(s.chatws)).Methods("GET")
	r.HandleFunc("/chat/notify", s.notify).Methods("GET")

	s.r = r
}

func (s *Server) swaggerDocs() {
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	docs.SwaggerInfo.Schemes = strings.Split(os.Getenv("SWAGGER_SCHEMES"), ",")
	docs.SwaggerInfo.BasePath = os.Getenv("SWAGGER_BASE_PATH")

	// Load swagger documentation HTML
	s.r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler).Methods("GET")

	// Load swagger documentation JSON
	s.r.HandleFunc("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, fmt.Sprintf("%s/swagger.json", os.Getenv("SWAGGER_DOCS_PATH")))
	}).Methods("GET")
}

func (s *Server) Run(port string) error {
	log.Printf("Starting server at port: %s", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), s.r)
}
