package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ObjectID string

func (id ObjectID) MarshalBSONValue() (bsontype.Type, []byte, error) {
    p, err := primitive.ObjectIDFromHex(string(id))
    if err != nil {
        return bsontype.Null, nil, err
    }
    return bson.MarshalValue(p)
}
