package mongodb

import (
	"context"
	"log"
	"sync"
	"sync/atomic"

	"github.com/tanveerprottoy/starter-go/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	instance    *Client
	once        sync.Once
	mu          sync.Mutex
	initialized uint32
)

type Client struct {
	DBClient *mongo.Client
	DB       *mongo.Database
}

func GetInstance() *Client {
	once.Do(func() {
		instance = new(Client)
		instance.connect()
	})
	return instance
}

func GetInstanceMutex() *Client {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = new(Client)
			instance.connect()
		}
	}
	return instance
}

func GetInstanceAtomic() *Client {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()
	if initialized == 0 {
		instance = new(Client)
		instance.connect()
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}

func (d *Client) connect() {
	// uri := config.GetEnvValue("DB_URI")
	uri := config.GetJsonValue("dbUri").(string)
	/* credential := options.Credential{
		Username: "username",
		Password: "pass",
	} */
	var err error
	ctx := context.TODO()
	// opts := options.Client().ApplyURI("mongodb+srv://<host>").SetAuth(credential)
	opts := options.Client().ApplyURI(uri)
	d.DBClient, err = mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}
	err = d.DBClient.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected!")
	d.DB = d.DBClient.Database(config.GetEnvValue("DB_NAME"))
	// Establish and verify connection
	err = d.DB.Client().Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected successfully to DB")
}

func (d *Client) Disconnect() {
	if err := d.DBClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
