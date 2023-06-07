package mongo

import (
	"context"

	"github.com/lordvidex/errs"

	mongodb "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(ctx context.Context, url string) (*mongodb.Client, error) {
	// create mongo client
	client, err := mongodb.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, errs.B(err).Code(errs.Unavailable).Msg("failed to connect to mongo server").Err()
	}
	// ping mongo server
	if err = client.Ping(ctx, nil); err != nil {
		return nil, errs.B(err).Code(errs.Unavailable).Msg("failed to ping mongo server").Err()
	}
	return client, nil
}
