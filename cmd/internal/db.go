package internal

import (
	"log"
	"os"
	"path"

	"github.com/TheBoringDude/minidb"
)

// new minidb instance wrapper
func DB() *minidb.MiniDB {
	confDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	return minidb.New(path.Join(confDir, "simple-store"))
}

// checks the key if exists and returns a MiniCollections instance
func GetCols(key string) *minidb.MiniCollections {
	iDb := DB()

	if _, err := iDb.FindCollection(key); err != nil {
		log.Fatalf("Collections key: `%s` does not exist!", key)
	}

	return iDb.Collections(key)
}

// checks the key if exists and returns a MiniStore instance
func GetStore(key string) *minidb.MiniStore {
	iDb := DB()

	if _, err := iDb.FindStore(key); err != nil {
		log.Fatalf("Stores key: `%s` does not exist!", key)
	}

	return iDb.Store(key)
}
