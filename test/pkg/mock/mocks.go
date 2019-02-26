package mock

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

type MockMongoDataStore struct {
	FindValue []interface{}
}

func (client MockMongoDataStore) Insert(db string, collectionName string, obj interface{}) (id primitive.ObjectID, err error) {
	return primitive.NewObjectID(), nil
}

func (client MockMongoDataStore) Find(db string, collectionName string, filter interface{}, opts ...*options.FindOptions) ([]interface{}, error) {
	return client.FindValue, nil
}

func (client MockMongoDataStore) InsertMany(db string, collectionName string, obj ...interface{}) (ids []primitive.ObjectID, err error) {
	for range obj {
		ids = append(ids, primitive.NewObjectID())
	}
	return ids, nil
}
