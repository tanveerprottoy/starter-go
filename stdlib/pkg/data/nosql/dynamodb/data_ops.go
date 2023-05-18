package dynamodb

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	instanceDataOps *ControlOps
	onceDataOps     sync.Once
)

type DataOps struct {
	DBClient *dynamodb.Client
}

func NewDataOps(c *dynamodb.Client) *ControlOps {
	onceControlOps.Do(func() {
		instanceControlOps = &ControlOps{DBClient: c}
	})
	return instanceControlOps
}

func (c *ControlOps) PutItem(i *dynamodb.PutItemInput, ctx context.Context) (*dynamodb.PutItemOutput, error) {
	// Build the request with its input parameters
	return c.DBClient.PutItem(ctx, i)
}

func (c *ControlOps) BatchWriteItem(i *dynamodb.BatchWriteItemInput, ctx context.Context) (*dynamodb.BatchWriteItemOutput, error) {
	// Build the request with its input parameters
	return c.DBClient.BatchWriteItem(ctx, i)
}

func (c *ControlOps) TransactWriteItems(i *dynamodb.TransactWriteItemsInput, ctx context.Context) (*dynamodb.TransactWriteItemsOutput, error) {
	return c.DBClient.TransactWriteItems(ctx, i)
}

func (c *ControlOps) Query(i *dynamodb.QueryInput, ctx context.Context) (*dynamodb.QueryOutput, error) {
	// Build the request with its input parameters
	return c.DBClient.Query(ctx, i)
}

func (c *ControlOps) Scan(i *dynamodb.ScanInput, ctx context.Context) (*dynamodb.ScanOutput, error) {
	return c.DBClient.Scan(ctx, i)
}

func (c *ControlOps) GetItem(i *dynamodb.GetItemInput, ctx context.Context) (*dynamodb.GetItemOutput, error) {
	return c.DBClient.GetItem(ctx, i)
}

func (c *ControlOps) BatchGetItem(i *dynamodb.BatchGetItemInput, ctx context.Context) (*dynamodb.BatchGetItemOutput, error) {
	return c.DBClient.BatchGetItem(ctx, i)
}

func (c *ControlOps) TransactGetItems(i *dynamodb.TransactGetItemsInput, ctx context.Context) (*dynamodb.TransactGetItemsOutput, error) {
	return c.DBClient.TransactGetItems(ctx, i)
}

func (c *ControlOps) UpdateItem(i *dynamodb.UpdateItemInput, ctx context.Context) (*dynamodb.UpdateItemOutput, error) {
	return c.DBClient.UpdateItem(ctx, i)
}

func (c *ControlOps) DeleteItem(i *dynamodb.DeleteItemInput, ctx context.Context) (*dynamodb.DeleteItemOutput, error) {
	return c.DBClient.DeleteItem(ctx, i)
}
