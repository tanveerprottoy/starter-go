package repository

import (
	"context"
	"txp/restapistarter/pkg/data/nosql/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserMongoRepository[T any] struct {
}

func (r *UserMongoRepository[T]) Create(
	collectionName string,
	ctx context.Context,
	doc T,
	opts ...*options.InsertOneOptions,
) (*mongo.InsertOneResult, error) {
	return nil, nil
}

func (r *UserMongoRepository[T]) ReadMany(
	collectionName string,
	ctx context.Context,
	filter any,
	opts ...*options.FindOptions,
) (*mongo.Cursor, error) {
	return mongodb.Find(
		collectionName,
		ctx,
		filter,
		opts...,
	)
}

func (r *UserMongoRepository[T]) ReadOne(
	collectionName string,
	ctx context.Context,
	filter any,
	opts ...*options.FindOneOptions,
) *mongo.SingleResult {
	return nil
}

func (r *UserMongoRepository[T]) Update(
	collectionName string,
	ctx context.Context,
	filter any,
	doc T,
	opts ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	return nil, nil
}

func (r *UserMongoRepository[T]) Delete(
	collectionName string,
	ctx context.Context,
	filter any,
	opts ...*options.DeleteOptions,
) (*mongo.DeleteResult, error) {
	return nil, nil
}
