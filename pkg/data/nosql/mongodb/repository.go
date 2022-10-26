package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository[T any] interface {
	Create(
		collectionName string,
		ctx context.Context,
		doc T,
		opts ...*options.InsertOneOptions,
	) (*mongo.InsertOneResult, error)

	ReadMany(
		collectionName string,
		ctx context.Context,
		filter any,
		opts ...*options.FindOptions,
	) (*mongo.Cursor, error)

	ReadOne(
		collectionName string,
		ctx context.Context,
		filter any,
		opts ...*options.FindOneOptions,
	) *mongo.SingleResult

	Update(
		collectionName string,
		ctx context.Context,
		filter any,
		doc T,
		opts ...*options.UpdateOptions,
	) (*mongo.UpdateResult, error)

	Delete(
		collectionName string,
		ctx context.Context,
		filter any,
		opts ...*options.DeleteOptions,
	) (*mongo.DeleteResult, error)
}
