package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Create(ctx context.Context, doc any, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)

	ReadMany(ctx context.Context, filter any, opts ...*options.FindOptions) (*mongo.Cursor, error)

	ReadOne(ctx context.Context, filter any, opts ...*options.FindOneOptions) *mongo.SingleResult

	Update(ctx context.Context, filter any, doc any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)

	Delete(ctx context.Context, filter any, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
}
