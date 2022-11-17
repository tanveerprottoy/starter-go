package mongodb

import (
	"context"
	"log"
	"sync"
	"txp/restapistarter/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBClient struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func NewDBClient() *DBClient {
	d := new(DBClient)
	var once sync.Once
	once.Do(func() {
		d.connect()
	})
	return d
}

func (d *DBClient) connect() {
	// uri := config.GetEnvValue("DB_URI")
	// println(uri)
	credential := options.Credential{
		Username: "username",
		Password: "pass",
	}
	var err error
	ctx := context.TODO()
	opts := options.Client().ApplyURI("mongodb+srv://<host>").SetAuth(credential)
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
