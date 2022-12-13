package repository

import (
	"context"
	"txp/restapistarter/internal/pkg/constant"
	"txp/restapistarter/pkg/data/nosql/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserMongoRepository struct {
	db *mongo.Database
}

func NewUserMongoRepository(db *mongo.Database) *UserMongoRepository {
	r := new(UserMongoRepository)
	r.db = db
	return r
}

func (r *UserMongoRepository) Create(
	ctx context.Context,
	doc any,
	opts ...*options.InsertOneOptions,
) (*mongo.InsertOneResult, error) {
	return mongodb.InsertOne(
		r.db,
		constant.UsersCollection,
		ctx,
		doc,
		opts...,
	)
}

func (r *UserMongoRepository) ReadMany(
	ctx context.Context,
	filter any,
	opts ...*options.FindOptions,
) (*mongo.Cursor, error) {
	return mongodb.Find(
		r.db,
		constant.UsersCollection,
		ctx,
		filter,
		opts...,
	)
}

func (r *UserMongoRepository) ReadOne(
	ctx context.Context,
	filter any,
	opts ...*options.FindOneOptions,
) *mongo.SingleResult {
	return mongodb.FindOne(
		r.db,
		constant.UsersCollection,
		ctx,
		filter,
		opts...,
	)
}

func (r *UserMongoRepository) Update(
	ctx context.Context,
	filter any,
	doc any,
	opts ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	return mongodb.UpdateOne(
		r.db,
		constant.UsersCollection,
		ctx,
		filter,
		doc,
		opts...,
	)
}

func (r *UserMongoRepository) Delete(
	ctx context.Context,
	filter any,
	opts ...*options.DeleteOptions,
) (*mongo.DeleteResult, error) {
	return mongodb.DeleteOne(
		r.db,
		constant.UsersCollection,
		ctx,
		filter,
		opts...,
	)
}
