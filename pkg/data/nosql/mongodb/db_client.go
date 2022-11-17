package mongodb

import (
	"context"
	"log"
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
	uri := config.GetEnvValue("DB_URI")
	var err error
	ctx := context.TODO()
	d.Client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
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
	return d
}

func (d *DBClient) Disconnect() {
	if err := d.Client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
