package dynamodb

import (
	"context"
	"log"
	"sync"
	"sync/atomic"

	configPkg "github.com/tanveerprottoy/starter-go/pkg/config"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	instanceClient    *Client
	onceClient        sync.Once
	muClient          sync.Mutex
	initializedClient uint32
)

type Client struct {
	DBClient *dynamodb.Client
}

func NewClient() *Client {
	once.Do(func() {
		instanceClient = new(Client)
		instanceClient.init()
	})
	return instanceClient
}

func NewClientMutex() *Client {
	if instanceClient == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instanceClient = new(Client)
			instanceClient.init()
		}
	}
	return instanceClient
}

func NewClientAtomic() *Client {
	if atomic.LoadUint32(&initialized) == 1 {
		return instanceClient
	}
	mu.Lock()
	defer mu.Unlock()
	if initialized == 0 {
		instanceClient = new(Client)
		instanceClient.init()
		atomic.StoreUint32(&initialized, 1)
	}
	return instanceClient
}

func (c *Client) init() {
	// uri := config.GetEnvValue("DB_URI")
	reg := configPkg.GetJsonValue("region").(string)
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(reg))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	// Using the Config value, create the DynamoDB client
	c.DBClient = dynamodb.NewFromConfig(cfg)
	log.Println("init db success")
}
