package gapi

import (
	"fmt"

	config "simplesedge.com/feed/pkg/config"
	"simplesedge.com/feed/pkg/db"
	"simplesedge.com/gokit/token"
	pb "simplesedge.com/proto/gen/go/webapis/v1alpha1"
)

// serves all gRPC requests for our feed service
type Server struct {
	pb.UnimplementedSimpleSedgeServer
	config     config.Config
	store      db.Store
	tokenMaker token.Maker
}

// Creates a new gRPC server.
func NewServer(config *config.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create a token maker: %w", err)
	}

	server := &Server{
		config:     *config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}