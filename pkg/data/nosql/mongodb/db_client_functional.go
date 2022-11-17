package mongodb

import (
	"context"
	"log"
	"txp/restapistarter/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	DB     *mongo.Database
)

func InitDBClient() {
	uri := config.GetEnvValue("DB_URI")
	var err error
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	log.Println("Successfully connected!")
	DB = Client.Database(config.GetEnvValue("DB_NAME"))
	// Establish and verify connection
	err = DB.Client().Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected successfully to DB")
}

func Disconnect() {
	if err := Client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
