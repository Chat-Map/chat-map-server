package mongo

import (
	"context"
	"fmt"

	mongodb "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(ctx context.Context, url string) (*mongodb.Client, error) {
	// create mongo client
	client, err := mongodb.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo server: %s", err)
	}
	// ping mongo server
	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongo server: %s", err)
	}
	return client, nil
}
