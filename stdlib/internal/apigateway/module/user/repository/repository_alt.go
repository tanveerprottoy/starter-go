package repository

import (
	"context"

	"github.com/tanveerprottoy/starter-go/stdlib/internal/pkg/constant"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/data/nosql/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RepositoryAlt struct {
	db *mongo.Database
}

func NewRepositoryAlt(db *mongo.Database) *RepositoryAlt {
	r := new(RepositoryAlt)
	r.db = db
	return r
}

func (r *RepositoryAlt) Create(
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

func (r *RepositoryAlt) ReadMany(
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

func (r *RepositoryAlt) ReadOne(
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

func (r *RepositoryAlt) Update(
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

func (r *RepositoryAlt) Delete(
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
