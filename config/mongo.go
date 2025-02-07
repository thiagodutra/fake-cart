package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectDatabase(uri string, maxPoolSize uint64) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri).SetMaxPoolSize(maxPoolSize)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		panic(fmt.Sprintf("failed to connect to the database: %v", err))
	}

	if err := client.Ping(ctx, nil); err != nil {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(fmt.Sprintf("failed to disconnect from database: %v", err))
		}
		return nil, fmt.Errorf("failed to disconnect from the database %w", err)
	}

	return client, nil
}
