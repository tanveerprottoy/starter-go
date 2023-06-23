package dynamodb

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/pkg/errors"
)

var (
	instanceControlOps *ControlOps
	onceControlOps     sync.Once
	mutControlOps      sync.Mutex
)

type ControlOps struct {
	DBClient *dynamodb.Client
}

func NewControlOps(c *dynamodb.Client) *ControlOps {
	onceControlOps.Do(func() {
		instanceControlOps = &ControlOps{DBClient: c}
	})
	return instanceControlOps
}

func (c *ControlOps) CreateTable(i *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
	// Build the request with its input parameters
	out, err := c.DBClient.CreateTable(context.TODO(), i)
	return out, err
}

func (c *ControlOps) ListTables() {
	// Build the request with its input parameters
	resp, err := c.DBClient.ListTables(context.TODO(), &dynamodb.ListTablesInput{
		Limit: aws.Int32(5),
	})
	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
	}
	fmt.Println("Tables:")
	for _, tableName := range resp.TableNames {
		fmt.Println(tableName)
	}
}

func waitForTable(ctx context.Context, db *dynamodb.Client, tn string) error {
    w := dynamodb.NewTableExistsWaiter(db)
    err := w.Wait(ctx,
        &dynamodb.DescribeTableInput{
            TableName: aws.String(tn),
        },
        2*time.Minute,
        func(o *dynamodb.TableExistsWaiterOptions) {
            o.MaxDelay = 5 * time.Second
            o.MinDelay = 5 * time.Second
        })
    if err != nil {
        return errors.Wrap(err, "timed out while waiting for table to become active")
    }

    return err
}
