package db

import (
	"database/sql"
	"os"
	"testing"

	kitsql "simplesedge.com/gokit/sql"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simplesedge?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	testDB = kitsql.GetDB(dbDriver, dbSource)
	testQueries = New(testDB)
	os.Exit(m.Run())
}
