package main

import (
	"context"
	"log"
	"os"

	"github.com/Chat-Map/chat-map-server/internal/adapters/db/postgres"
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

	// Run migrations
	err = postgres.Migrate(pg, os.Getenv("MIGRATION_DIR"))
	if err != nil {
		log.Fatalf("error running migrations: %+v", err)
	}

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

	// Create usecases
	uc := application.NewUseCase(
		application.WithUserRepository(ur),
		application.WithSessionsRepository(sr),
		application.WithChatRepository(cr),
		application.WithMessageRepository(mr),
	)

	_ = uc
	_ = appCtx
}
