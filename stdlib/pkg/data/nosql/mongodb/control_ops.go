package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ListDatabases(c *mongo.Client) (mongo.ListDatabasesResult, error) {
	return c.ListDatabases(context.Background(), nil)
}

func ListDatabaseNames(c *mongo.Client) ([]string, error) {
	return c.ListDatabaseNames(context.Background(), nil)
}

func GetCollection(
	db *mongo.Database,
	name string,
	opts ...*options.CollectionOptions,
) *mongo.Collection {
	return db.Collection(name, opts...)
}
