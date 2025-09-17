// Package mongodb contains logic to connect to MongoDB.
package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBClient is a MongoDB client
type MongoDBClient struct {
	client     *mongo.Client
	collection *mongo.Collection
}

const (
	databaseName = "user"
)

// NewMongoDBClient creates a new MongoDB client
func NewMongoDBClient(ctx context.Context, uri string, collectionName string) (*MongoDBClient, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	database := client.Database(databaseName)
	collection := database.Collection(collectionName)
	return &MongoDBClient{client: client, collection: collection}, nil
}

// Close closes the MongoDB client
func (c *MongoDBClient) Close(ctx context.Context) error {
	return c.client.Disconnect(ctx)
}

// GetClient returns the MongoDB client
func (c *MongoDBClient) GetClient() *mongo.Client {
	return c.client
}

// GetCollection returns the MongoDB collection
func (c *MongoDBClient) GetCollection() *mongo.Collection {
	return c.collection
}
