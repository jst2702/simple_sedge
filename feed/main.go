package main

import (
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"simplesedge.com/feed/api"
	"simplesedge.com/feed/gapi"
	config "simplesedge.com/feed/pkg/config"
	"simplesedge.com/feed/pkg/db"
	kitsql "simplesedge.com/gokit/sql"
	pb "simplesedge.com/proto/gen/go/webapis/v1alpha1"
)

func main() {
	cfg := config.DefaultConfig(".")
	conn := kitsql.GetDB(cfg.DBDriver, cfg.DBSource)
	store := db.NewStore(conn)
	runGRPCServer(cfg, store)
	// runGinServer(cfg, store)
}

func runGRPCServer(cfg *config.Config, store db.Store) {
	server, err := gapi.NewServer(cfg, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleSedgeServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", cfg.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}

}

func runGinServer(cfg *config.Config, store db.Store) {
	server, err := api.NewServer(cfg, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(cfg.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
