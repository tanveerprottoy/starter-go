package mongodb

import (
	"context"
	"log"
	"txp/restapistarter/pkg/configutil"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
)

func InitDBClient() {
	uri := configutil.GetEnvValue("MONGODB_URI")
	var err error
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	log.Println("Successfully connected!")
}
