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
