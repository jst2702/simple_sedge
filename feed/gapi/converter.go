package gapi

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"simplesedge.com/feed/pkg/db"
	pb "simplesedge.com/proto/gen/go/webapis/v1alpha1"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
