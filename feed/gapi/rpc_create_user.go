package gapi

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"simplesedge.com/feed/pkg/db"
	"simplesedge.com/gokit/util"
	pb "simplesedge.com/proto/gen/go/webapis/v1alpha1"
)

func (server *Server) CreateUser(
	ctx context.Context,
	req *pb.CreateUserRequest,
) (*pb.CreateUserResponse, error) {

	if err := server.validEmail(req.Email); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid email %s", err)
	}

	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password %s", err)
	}

	arg := db.CreateUserParams{
		Email:          req.GetEmail(),
		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username or email already exists %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create user %s", err)
	}

	rsp := &pb.CreateUserResponse{
		User: convertUser(user),
	}

	return rsp, nil
}

func (server *Server) validEmail(email string) error {
	if len(email) < 10 {
		err := fmt.Errorf("email is too short, desu, %d", 10)
		return err
	} else if strings.Contains(email, "yahoo.com") {
		err := errors.New("yahoo isn't a real email")
		return err
	} else {
		return nil
	}
}
