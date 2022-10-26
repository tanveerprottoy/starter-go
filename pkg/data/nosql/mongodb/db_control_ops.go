package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ListDatabases() (mongo.ListDatabasesResult, error) {
	return Client.ListDatabases(context.Background(), nil)
}

func ListDatabaseNames() ([]string, error) {
	return Client.ListDatabaseNames(context.Background(), nil)
}

func GetCollection(
	name string,
	opts ...*options.CollectionOptions,
) *mongo.Collection {
	return DB.Collection(name, opts...)
}
