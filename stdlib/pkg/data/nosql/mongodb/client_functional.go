package mongodb

import (
	"context"
	"log"

	"github.com/tanveerprottoy/starter-go/stdlib/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DBClient *mongo.Client
	DB       *mongo.Database
)

func InitDBClient() {
	uri := config.GetEnvValue("DB_URI")
	var err error
	DBClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	log.Println("Successfully connected!")
	DB = DBClient.Database(config.GetEnvValue("DB_NAME"))
	// Establish and verify connection
	err = DB.Client().Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected successfully to DB")
}

func Disconnect() {
	if err := DBClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
