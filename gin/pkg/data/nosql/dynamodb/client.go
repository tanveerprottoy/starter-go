package dynamodb

import (
	"context"
	"log"
	"sync"
	"sync/atomic"

	configPkg "github.com/tanveerprottoy/starter-go/stdlib/pkg/config"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	instanceClient    *Client
	onceClient        sync.Once
	mutClient         sync.Mutex
	initializedClient uint32
)

type Client struct {
	DBClient *dynamodb.Client
}

func NewClient() *Client {
	onceClient.Do(func() {
		instanceClient = new(Client)
		instanceClient.init()
	})
	return instanceClient
}

func NewClientMutex() *Client {
	if instanceClient == nil {
		mutClient.Lock()
		defer mutClient.Unlock()
		if instanceClient == nil {
			instanceClient = new(Client)
			instanceClient.init()
		}
	}
	return instanceClient
}

func NewClientAtomic() *Client {
	if atomic.LoadUint32(&initializedClient) == 1 {
		return instanceClient
	}
	mutClient.Lock()
	defer mutClient.Unlock()
	if initializedClient == 0 {
		instanceClient = new(Client)
		instanceClient.init()
		atomic.StoreUint32(&initializedClient, 1)
	}
	return instanceClient
}

func (c *Client) init() {
	// uri := config.GetEnvValue("DB_URI")
	reg := configPkg.GetJsonValue("region").(string)
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(reg))
	/* cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("AKID", "SECRET_KEY", "TOKEN")),
		)
		cfg, err := config.LoadDefaultConfig(context.Background(), func(o *config.LoadOptions) error {
	        o.Region = reg
	        return nil
	    })
		/* cfg , err := config.LoadDefaultConfig(context.TODO(),
	    config.WithSharedCredentialsFiles(
		[]string{"test/credentials", "data/credentials"},
	    ),
	    config.WithSharedConfigFiles(
	        []string{"test/config", "data/config"},
		) */
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	// Using the Config value, create the DynamoDB client
	c.DBClient = dynamodb.NewFromConfig(cfg)
	log.Println("init db success")
}
