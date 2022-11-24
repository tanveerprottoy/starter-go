package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DecodeCursor[T any](c *mongo.Cursor, ctx context.Context) (T, error) {
	var data T
	err := c.All(ctx, &data)
	return data, err
}

func DecodeSingleResult[T any](c *mongo.SingleResult) (T, error) {
	var datum T
	err := c.Decode(&datum)
	return datum, err
}

func BuildObjectID(id string) (primitive.ObjectID, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	return objId, err
}

func BuildPaginatedOpts(limit, skip int) options.FindOptions {
	l := int64(limit)
	s := int64(skip*limit - limit)
	return options.FindOptions{Limit: &l, Skip: &s}
}
