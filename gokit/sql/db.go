package sql

import (
	"database/sql"
	"log"
)

func GetDB(
	dbDriver string,
	dbSource string) (
	db *sql.DB) {

	var err error

	db, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db.", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("unsuccessful ping of db.", err)
	}

	return db
}
