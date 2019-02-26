package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

type DataStore interface {
	Insert(db string, collectionName string, obj interface{}) (id primitive.ObjectID, err error)
	Find(db string, collectionName string, filter interface{}, opts ...options.FindOptions) (*mongo.Cursor, error)
}

type MongoDataStore struct {
	*mongo.Client
}

func NewDbConnection(config *Config) (MongoDataStore, error){
	mdbFullUrl := fmt.Sprintf("mongodb+srv://%s:%s@%s", config.MdbUser, config.MdbPassword, config.MdbUrl)
	client, err := mongo.Connect(context.TODO(), mdbFullUrl)
	if err == nil {
		err = client.Ping(context.TODO(), nil)
	}
	return MongoDataStore{client}, err
}

func (client MongoDataStore) Insert(db string, collectionName string, obj interface{}) (id primitive.ObjectID, err error) {
	collection := client.Database(db).Collection(collectionName)
	result, err := collection.InsertOne(context.TODO(), &obj)
	if id, ok := result.InsertedID.(primitive.ObjectID); ok {
		return id, nil
	}
	return id, errors.New("unable to get object ID from inserted item")
}

func (client MongoDataStore) Find(db string, collectionName string, filter interface{}, opts ...options.FindOptions) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetLimit(2)

	collection := client.Database(db).Collection(collectionName)
	return collection.Find(context.TODO(), filter, findOptions)
}
