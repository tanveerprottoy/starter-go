package repository

import (
	"context"
	"txp/restapistarter/pkg/data/nosql/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserMongoRepository struct {
	DB *mongo.Database
}

func NewUserMongoRepository(db *mongo.Database) *UserMongoRepository {
	r := new(UserMongoRepository)
	r.DB = db
	return r
}

func (r *UserMongoRepository) Create(
	collectionName string,
	ctx context.Context,
	doc any,
	opts ...*options.InsertOneOptions,
) (*mongo.InsertOneResult, error) {
	return mongodb.InsertOne(
		r.DB,
		collectionName,
		ctx,
		doc,
		opts...,
	)
}

func (r *UserMongoRepository) ReadMany(
	collectionName string,
	ctx context.Context,
	filter any,
	opts ...*options.FindOptions,
) (*mongo.Cursor, error) {
	return mongodb.Find(
		r.DB,
		collectionName,
		ctx,
		filter,
		opts...,
	)
}

func (r *UserMongoRepository) ReadOne(
	collectionName string,
	ctx context.Context,
	filter any,
	opts ...*options.FindOneOptions,
) *mongo.SingleResult {
	return mongodb.FindOne(
		r.DB,
		collectionName,
		ctx,
		filter,
		opts...,
	)
}

func (r *UserMongoRepository) Update(
	collectionName string,
	ctx context.Context,
	filter any,
	doc any,
	opts ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	return mongodb.UpdateOne(
		r.DB,
		collectionName,
		ctx,
		filter,
		doc,
		opts...,
	)
}

func (r *UserMongoRepository) Delete(
	collectionName string,
	ctx context.Context,
	filter any,
	opts ...*options.DeleteOptions,
) (*mongo.DeleteResult, error) {
	return mongodb.DeleteOne(
		r.DB,
		collectionName,
		ctx,
		filter,
		opts...,
	)
}
