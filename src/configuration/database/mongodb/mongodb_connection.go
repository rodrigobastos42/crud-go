package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDBConnection(
	ctx context.Context,
) (*mongo.Database, error) {
	mongoDBUri := os.Getenv("MONGODB_URI")
	mongoDBDatabase := os.Getenv("MONGODB_DATABASE")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDBUri))

	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(mongoDBDatabase), nil
}
