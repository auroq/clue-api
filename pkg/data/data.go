package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"net/url"
	"os"
)

type dataStore interface {
	insert(db string, collectionName string, obj interface{}) (primitive.ObjectID, error)
	find(db string, collectionName string, filter interface{}, opts ...options.FindOptions) (*mongo.Cursor, error)
	disconnect() error
}

type mongoDataStore struct {
	*mongo.Client
}

func (client mongoDataStore) insert(db string, collectionName string, obj interface{}) (id primitive.ObjectID, err error) {
	collection := client.Database(db).Collection(collectionName)
	result, err := collection.InsertOne(context.TODO(), &obj)
	if id, ok := result.InsertedID.(primitive.ObjectID); ok {
		return id, nil
	}
	return id, errors.New("unable to get object ID from inserted item")
}

func (client mongoDataStore) find(db string, collectionName string, filter interface{}, opts ...options.FindOptions) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetLimit(2)

	collection := client.Database(db).Collection(collectionName)
	return collection.Find(context.TODO(), filter, findOptions)
}

func (client mongoDataStore) disconnect() error {
	return client.Disconnect(context.TODO())
}

func getConnection() (dataStore, error) {
	mdbUrl := os.Getenv("CLUE_MDB_URL")
	mdbUser := os.Getenv("CLUE_MDB_USER")
	mdbPass := os.Getenv("CLUE_MDB_PASSWORD")
	mdbPass = url.QueryEscape(mdbPass)
	mdbFullUrl := fmt.Sprintf("mongodb+srv://%s:%s@%s", mdbUser, mdbPass, mdbUrl)

	client, err := mongo.Connect(context.TODO(), mdbFullUrl)
	if err == nil {
		err = client.Ping(context.TODO(), nil)
	}
	return mongoDataStore{client}, err
}
