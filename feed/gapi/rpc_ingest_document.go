package gapi

import (
	"context"
	"database/sql"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"simplesedge.com/feed/pkg/db"
	"simplesedge.com/feed/val"
	ppb "simplesedge.com/proto/gen/go/processing/v1alpha1"
)

func (server *Server) IngestDocument(
	ctx context.Context,
	req *ppb.IngestDocumentRequest,
) (*ppb.IngestDocumentResponse, error) {
	apiKey, err := server.authorizeApiKey(ctx, req.ApiKey)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateIngestDocumentRequest(req)
	if violations != nil {
		return nil, InvalidArgumentError(violations)
	}

	ticker := req.GetTicker()
	language := req.GetLanguage()

	arg := db.CreateDocumentParams{
		Guid:     req.GetGuid(),
		Url:      req.GetUrl(),
		Site:     req.GetSite(),
		SiteFull: req.GetSiteFull(),
		Headline: req.GetHeadline(),
		Title:    req.GetTitle(),
		Body:     req.GetBody(),
		Ticker: sql.NullString{
			String: ticker,
			Valid:  ticker == "",
		},
		Tickers:     req.GetTickers(),
		PublishedAt: req.GetPublishedAt().AsTime(),
		Language: sql.NullString{
			String: language,
			Valid:  language == "",
		},
		ApiKeyUsed: apiKey.ApiKey,
	}

	document, err := server.store.CreateDocument(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create document %s", err)
	}

	rsp := &ppb.IngestDocumentResponse{
		Document: convertDocument(document),
	}

	return rsp, nil
}

func validateIngestDocumentRequest(
	req *ppb.IngestDocumentRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if req.GetTicker() != "" {
		if err := val.ValidateTicker(req.GetTicker(), true); err != nil {
			violations = append(violations, fieldViolation("ticker", err))
		}
	}
	for _, ticker := range req.GetTickers() {
		if err := val.ValidateTicker(ticker, false); err != nil {
			violations = append(violations, fieldViolation("tickers", err))
		}
	}

	// Maybe add other checks for params
	return violations

}
