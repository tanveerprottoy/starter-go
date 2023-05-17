package dynamodb

import (
	"context"
	"fmt"
	"log"
	"sync"

	configPkg "github.com/tanveerprottoy/starter-go/pkg/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	instanceControlOps    *ControlOps
	onceControlOps        sync.Once
	muControlOps          sync.Mutex
	initializedControlOps uint32
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

func (c *ControlOps) CreateTable() {
	// Build the request with its input parameters
	resp, err := c.DBClient.CreateTable(context.TODO(), &dynamodb.ListTablesInput{
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
