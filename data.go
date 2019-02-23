package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"net/url"
	"os"
)


func connectToMongoDB() {
	mdbUrl := os.Getenv("CLUE_MDB_URL")
	mdbUser := os.Getenv("CLUE_MDB_USER")
	mdbPass := os.Getenv("CLUE_MDB_PASSWORD")
	mdbPass = url.QueryEscape(mdbPass)
	mdbFullUrl := fmt.Sprintf("mongodb+srv://%s:%s@%s", mdbUser, mdbPass, mdbUrl)

	client, err := mongo.Connect(context.TODO(), mdbFullUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}