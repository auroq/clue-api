package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"net/url"
	"os"
)

type dataStore interface {
	insert(db string, collectionName string, player Player) (Player, error)
	disconnect() error
}

type mongoDataStore struct {
	client *mongo.Client
}

func (store mongoDataStore) insert(db string, collectionName string, player Player) (Player, error) {
	client := store.client
	collection := client.Database(db).Collection(collectionName)
	id, err := collection.InsertOne(context.TODO(), &player)
	player.ID = id
	return player, err
}

func (store mongoDataStore) disconnect() error {
	client := store.client
	return client.Disconnect(context.TODO())
}

func getConnection() (dataStore, error) {
	mdbUrl := os.Getenv("CLUE_MDB_URL")
	mdbUser := os.Getenv("CLUE_MDB_USER")
	mdbPass := os.Getenv("CLUE_MDB_PASSWORD")
	mdbPass = url.QueryEscape(mdbPass)
	mdbFullUrl := fmt.Sprintf("mongodb+srv://%s:%s@%s", mdbUser, mdbPass, mdbUrl)

	client, err := mongo.Connect(context.TODO(), mdbFullUrl)
	return mongoDataStore{client}, err
}