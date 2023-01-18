package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertOne(
	db *mongo.Database,
	collectionName string,
	ctx context.Context,
	doc any,
	opts ...*options.InsertOneOptions,
) (*mongo.InsertOneResult, error) {
	return db.Collection(collectionName).InsertOne(
		ctx,
		doc,
		opts...,
	)
}

func InsertMany(
	db *mongo.Database,
	collectionName string,
	ctx context.Context,
	docs []any,
	opts ...*options.InsertManyOptions,
) (*mongo.InsertManyResult, error) {
	return db.Collection(collectionName).InsertMany(
		ctx,
		docs,
		opts...,
	)
}

func Find(
	db *mongo.Database,
	collectionName string,
	ctx context.Context,
	filter any,
	opts ...*options.FindOptions,
) (*mongo.Cursor, error) {
	return db.Collection(collectionName).Find(
		ctx,
		filter,
		opts...,
	)
}

func FindOne(
	db *mongo.Database,
	collectionName string,
	ctx context.Context,
	filter any,
	opts ...*options.FindOneOptions,
) *mongo.SingleResult {
	return db.Collection(collectionName).FindOne(
		ctx,
		filter,
		opts...,
	)
}

func FindOneAndUpdate(
	db *mongo.Database,
	collectionName string,
	ctx context.Context,
	filter any,
	doc any,
	opts ...*options.FindOneAndUpdateOptions,
) *mongo.SingleResult {
	return db.Collection(collectionName).FindOneAndUpdate(
		ctx,
		filter,
		doc,
		opts...,
	)
}

func FindOneAndReplace(
	db *mongo.Database,
	collectionName string,
	ctx context.Context,
	filter any,
	doc any,
	opts ...*options.FindOneAndReplaceOptions,
) *mongo.SingleResult {
	return db.Collection(collectionName).FindOneAndReplace(
		ctx,
		filter,
		doc,
		opts...,
	)
}

func FindOneAndDelete(
	db *mongo.Database,
	collectionName string,
	ctx context.Context,
	filter any,
	opts ...*options.FindOneAndDeleteOptions,
) *mongo.SingleResult {
	return db.Collection(collectionName).FindOneAndDelete(
		ctx,
		filter,
		opts...,
	)
}

func UpdateMany(
	db *mongo.Database,
	collectionName string,
	ctx context.Context,
	filter any,
	docs any,
	opts ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	return db.Collection(collectionName).UpdateMany(
		ctx,
		filter,
		docs,
		opts...,
	)
}

func UpdateOne(
	db *mongo.Database,
	collectionName string,
	ctx context.Context,
	filter any,
	doc any,
	opts ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	return db.Collection(collectionName).UpdateOne(
		ctx,
		filter,
		doc,
		opts...,
	)
}

func UpdateByID(
	db *mongo.Database,
	collectionName string,
	ctx context.Context,
	filter any,
	doc any,
	opts ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	return db.Collection(collectionName).UpdateByID(
		ctx,
		filter,
		doc,
		opts...,
	)
}

func DeleteMany(
	db *mongo.Database,
	collectionName string,
	ctx context.Context,
	filter any,
	opts ...*options.DeleteOptions,
) (*mongo.DeleteResult, error) {
	return db.Collection(collectionName).DeleteMany(
		ctx,
		filter,
		opts...,
	)
}

func DeleteOne(
	db *mongo.Database,
	collectionName string,
	ctx context.Context,
	filter any,
	opts ...*options.DeleteOptions,
) (*mongo.DeleteResult, error) {
	return db.Collection(collectionName).DeleteOne(
		ctx,
		filter,
		opts...,
	)
}
