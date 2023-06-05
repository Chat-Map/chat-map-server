package main

import (
	"context"
	"log"
	"os"

	"github.com/Chat-Map/chat-map-server/internal/adapters/bcrypt"
	"github.com/Chat-Map/chat-map-server/internal/adapters/db/postgres"
	"github.com/Chat-Map/chat-map-server/internal/adapters/token/paseto"
	"github.com/Chat-Map/chat-map-server/internal/adapters/validator"
	"github.com/Chat-Map/chat-map-server/internal/api"
	"github.com/Chat-Map/chat-map-server/internal/application"
)

func main() {
	appCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create main db client
	pg, err := postgres.New(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("error creating postgres client: %+v", err)
	}
	log.Printf("postgres client created")

	// Run migrations
	err = postgres.Migrate(pg, os.Getenv("DATABASE_MIGRATION_PATH"))
	if err != nil {
		log.Fatalf("error running migrations: %+v", err)
	}
	log.Printf("migrations done")

	// Defer close postgres connection
	defer func() {
		err = pg.Close()
		if err != nil {
			log.Printf("error disconnecting from postgres: %+v", err)
		}
		log.Println("postgres client closed")
	}()

	// Create main db repositories
	ur := postgres.NewUserRepository(pg)
	sr := postgres.NewSessionRepository(pg)
	cr := postgres.NewChatRepository(pg)
	mr := postgres.NewMessageRepository(pg)

	// Create tokenizer
	tk, err := paseto.NewPaseto([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		log.Fatalf("error creating paseto: %+v", err)
	}
	log.Printf("paseto tokenizer created")

	// Create password hasher
	ph := bcrypt.New()

	// Create struct validator
	v := validator.New()

	// Create usecases
	uc := application.NewUseCase(
		application.WithValidator(v),
		application.WithTokenizer(tk),
		application.WithPasswordHasher(ph),

		application.WithUserRepository(ur),
		application.WithSessionsRepository(sr),
		application.WithChatRepository(cr),
		application.WithMessageRepository(mr),
	)

	server := api.NewServer(appCtx, uc)
	if err = server.Run(os.Getenv("APP_PORT")); err != nil {
		log.Fatalf("error running server: %+v", err)
	}
}
