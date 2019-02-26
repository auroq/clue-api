package data

import (
	"net/url"
	"os"
)

type Config struct {
	MdbUrl      string
	MdbUser     string
	MdbPassword string
}

func NewConfiguration() *Config {
	mdbUrl := os.Getenv("CLUE_MDB_URL")
	mdbUser := os.Getenv("CLUE_MDB_USER")
	mdbPass := os.Getenv("CLUE_MDB_PASSWORD")
	mdbPass = url.QueryEscape(mdbPass)
	return &Config{
		MdbUrl:      mdbUrl,
		MdbUser:     mdbUser,
		MdbPassword: mdbPass,
	}
}
