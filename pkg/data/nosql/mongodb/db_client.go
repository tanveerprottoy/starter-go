package mongodb

import (
	"context"
	"log"
	"sync"
	"sync/atomic"
	"txp/restapistarter/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	instance    *DBClient
	once        sync.Once
	mu          sync.Mutex
	initialized uint32
)

type DBClient struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func GetInstance() *DBClient {
	once.Do(func() {
		instance = new(DBClient)
		instance.connect()
	})
	return instance
}

func GetInstanceMutex() *DBClient {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = new(DBClient)
			instance.connect()
		}
	}
	return instance
}

func GetInstanceAtomic() *DBClient {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()
	if initialized == 0 {
		instance = new(DBClient)
		instance.connect()
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}

func (d *DBClient) connect() {
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
	d.Client, err = mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}
	err = d.Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected!")
	d.DB = d.Client.Database(config.GetEnvValue("DB_NAME"))
	// Establish and verify connection
	err = d.DB.Client().Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected successfully to DB")
}

func (d *DBClient) Disconnect() {
	if err := d.Client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
