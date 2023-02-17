package main

import (
	"log"

	_ "github.com/lib/pq"
	"simplesedge.com/feed/api"
	config "simplesedge.com/feed/pkg/config"
	"simplesedge.com/feed/pkg/db"
	kitsql "simplesedge.com/gokit/sql"
)

func main() {
	cfg := config.DefaultConfig(".")
	conn := kitsql.GetDB(cfg.DBDriver, cfg.DBSource)

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err := server.Start(cfg.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
