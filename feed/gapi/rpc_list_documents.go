package gapi

import (
	"context"
	"database/sql"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"simplesedge.com/feed/pkg/db"
	"simplesedge.com/feed/val"
	pb "simplesedge.com/proto/gen/go/apis/v1alpha1"
)

func (server *Server) ListDocuments(
	ctx context.Context,
	req *pb.ListDocumentsRequest,
) (*pb.ListDocumentsResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateListDocumentsRequest(req)
	if violations != nil {
		return nil, InvalidArgumentError(violations)
	}

	arg := db.ListDocumentsParams{
		Limit:  req.GetPageSize(),
		Offset: (req.GetPageID() - 1) * req.GetPageSize(),
	}

	documents, err := server.store.ListDocuments(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "documents not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to retrieve documents %s", err)
	}

	rsp := &pb.ListDocumentsResponse{
		Documents: convertDocuments(documents),
	}

	return rsp, nil
}

func validateListDocumentsRequest(
	req *pb.ListDocumentsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidatePageID(req.GetPageID()); err != nil {
		violations = append(violations, fieldViolation("pageId", err))
	}

	if err := val.ValidatePageSize(req.GetPageSize()); err != nil {
		violations = append(violations, fieldViolation("pageSize", err))
	}
	return violations

}
