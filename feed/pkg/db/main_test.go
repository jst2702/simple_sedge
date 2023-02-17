package db

import (
	"database/sql"
	"os"
	"testing"

	config "simplesedge.com/feed/pkg/config"
	kitsql "simplesedge.com/gokit/sql"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	cfg := config.DefaultConfig("../../")
	testDB = kitsql.GetDB(cfg.DBDriver, cfg.DBSource)
	testQueries = New(testDB)
	os.Exit(m.Run())
}
