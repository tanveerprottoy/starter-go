package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertOne(
	collectionName string,
	ctx context.Context,
	doc any,
	opts ...*options.InsertOneOptions,
) (*mongo.InsertOneResult, error) {
	return DB.Collection(collectionName).InsertOne(
		ctx,
		doc,
		opts...,
	)
}

func InsertMany(
	collectionName string,
	ctx context.Context,
	docs []any,
	opts ...*options.InsertManyOptions,
) (*mongo.InsertManyResult, error) {
	return DB.Collection(collectionName).InsertMany(
		ctx,
		docs,
		opts...,
	)
}

func Find(
	collectionName string,
	ctx context.Context,
	filter any,
	opts ...*options.FindOptions,
) (*mongo.Cursor, error) {
	return DB.Collection(collectionName).Find(
		ctx,
		filter,
		opts...,
	)
}

func FindOne(
	collectionName string,
	ctx context.Context,
	filter any,
	opts ...*options.FindOneOptions,
) *mongo.SingleResult {
	return DB.Collection(collectionName).FindOne(
		ctx,
		filter,
		opts...,
	)
}

func FindOneAndUpdate(
	collectionName string,
	ctx context.Context,
	filter any,
	doc any,
	opts ...*options.FindOneAndUpdateOptions,
) *mongo.SingleResult {
	return DB.Collection(collectionName).FindOneAndUpdate(
		ctx,
		filter,
		doc,
		opts...,
	)
}

func FindOneAndReplace(
	collectionName string,
	ctx context.Context,
	filter any,
	doc any,
	opts ...*options.FindOneAndReplaceOptions,
) *mongo.SingleResult {
	return DB.Collection(collectionName).FindOneAndReplace(
		ctx,
		filter,
		doc,
		opts...,
	)
}

func FindOneAndDelete(
	collectionName string,
	ctx context.Context,
	filter any,
	opts ...*options.FindOneAndDeleteOptions,
) *mongo.SingleResult {
	return DB.Collection(collectionName).FindOneAndDelete(
		ctx,
		filter,
		opts...,
	)
}

func UpdateMany(
	collectionName string,
	ctx context.Context,
	filter any,
	docs any,
	opts ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	return DB.Collection(collectionName).UpdateMany(
		ctx,
		filter,
		docs,
		opts...,
	)
}

func UpdateOne(
	collectionName string,
	ctx context.Context,
	filter any,
	doc any,
	opts ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	return DB.Collection(collectionName).UpdateOne(
		ctx,
		filter,
		doc,
		opts...,
	)
}

func UpdateByID(
	collectionName string,
	ctx context.Context,
	filter any,
	doc any,
	opts ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	return DB.Collection(collectionName).UpdateByID(
		ctx,
		filter,
		doc,
		opts...,
	)
}

func DeleteMany(
	collectionName string,
	ctx context.Context,
	filter any,
	opts ...*options.DeleteOptions,
) (*mongo.DeleteResult, error) {
	return DB.Collection(collectionName).DeleteMany(
		ctx,
		filter,
		opts...,
	)
}

func DeleteOne(
	collectionName string,
	ctx context.Context,
	filter any,
	opts ...*options.DeleteOptions,
) (*mongo.DeleteResult, error) {
	return DB.Collection(collectionName).DeleteOne(
		ctx,
		filter,
		opts...,
	)
}
