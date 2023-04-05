package gapi

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"simplesedge.com/feed/pkg/db"
	pb "simplesedge.com/proto/gen/go/apis/v1alpha1"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}

func convertDocument(doc db.Document) *pb.Document {
	return &pb.Document{
		Guid:        doc.Guid,
		Url:         doc.Url,
		Site:        doc.Site,
		SiteFull:    doc.SiteFull,
		SiteSection: doc.SiteSection,
		Headline:    doc.Headline,
		Title:       doc.Title,
		Body:        doc.Body,
		Tickers:     doc.Tickers,
		Language:    doc.Language.String,
	}
}

func convertDocuments(docs []db.Document) []*pb.Document {
	pb_docs := []*pb.Document{}
	for _, doc := range docs {
		pb_doc := convertDocument(doc)
		pb_docs = append(pb_docs, pb_doc)
	}
	return pb_docs
}
