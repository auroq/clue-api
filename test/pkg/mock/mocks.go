package mock

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"golang.org/x/crypto/openpgp/errors"
)

type MockMongoDataStore struct {
	data map[string]map[string]interface{}
}

func (client MockMongoDataStore) Insert(db string, collectionName string, obj interface{}) (id primitive.ObjectID, err error) {
	return primitive.NewObjectID(), nil
}

func (client MockMongoDataStore) Find(db string, collectionName string, filter interface{}, opts ...options.FindOptions) (*mongo.Cursor, error) {
	return nil, errors.UnsupportedError("function is not yet implemented")
}
