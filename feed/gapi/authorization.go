package gapi

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"
	"simplesedge.com/feed/pkg/db"
	"simplesedge.com/gokit/token"
)

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
)

func (server *Server) authorizeUser(ctx context.Context) (*token.Payload, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	values := md.Get(authorizationHeader)
	if len(values) == 0 {
		return nil, fmt.Errorf("missing authorization header")
	}

	authHeader := values[0]
	// <Authorization-type> <authorization data>
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return nil, fmt.Errorf("unsupported authorization type: %s", authType)
	}

	accessToken := fields[1]
	payload, err := server.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %s", err)
	}

	return payload, nil

}

func (server *Server) authorizeApiKey(ctx context.Context, keyString string) (db.ApiKey, error) {
	apiKey, err := server.store.GetApiKey(ctx, keyString)
	if err != nil {
		if err == sql.ErrNoRows {
			return apiKey, fmt.Errorf("api key not found")
		} else {
			return apiKey, fmt.Errorf("failed to retrieve api key: %w", err)
		}
	} else if !apiKey.Active {
		return apiKey, fmt.Errorf("inactive api key")
	}
	return apiKey, nil
}
