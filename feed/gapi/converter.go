package gapi

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"simplesedge.com/feed/pkg/db"
	ppb "simplesedge.com/proto/gen/go/processing/v1alpha1"
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

func convertDocument(doc db.Document) *ppb.Document {
	return &ppb.Document{
		Guid:        doc.Guid,
		Url:         doc.Url,
		Site:        doc.Site,
		SiteFull:    doc.SiteFull,
		SiteSection: doc.SiteSection,
		Headline:    doc.Headline,
		Title:       doc.Title,
		Body:        doc.Body,
		Ticker:      doc.Ticker.String,
		Tickers:     doc.Tickers,
		Language:    doc.Language.String,
	}
}

func convertDocuments(docs []db.Document) []*ppb.Document {
	ppb_docs := []*ppb.Document{}
	for _, doc := range docs {
		ppb_doc := convertDocument(doc)
		ppb_docs = append(ppb_docs, ppb_doc)
	}
	return ppb_docs
}
