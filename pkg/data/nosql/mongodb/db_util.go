package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Decode(c *mongo.Cursor) (bson.D, error) {
	var result bson.D
	err := c.Decode(&result)
	return result, err
}
