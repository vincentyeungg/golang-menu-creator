package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/vincentyeungg/golang-menu-creator/config"
)

// need to access queries for sql db calls
var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := config.LoadConfig("../..")
	conn, err := sql.Open(config.DbDriver, config.DbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
