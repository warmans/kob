package service

import "context"
import "github.com/warmans/kob/server/pkg/rpc"
import "github.com/golang/protobuf/ptypes/empty"

func NewKobService() *KobService {
	return &KobService{}
}

type KobService struct {
}

func (s *KobService) ListEntries(context.Context, *rpc.ListEntriesRequest) (*rpc.EntryList, error) {
	return &rpc.EntryList{Entries: []*rpc.Entry{{Title: "bar"}}}, nil
}

func (s *KobService) GetEntry(context.Context, *rpc.GetEntryRequest) (*rpc.Entry, error) {
	return &rpc.Entry{Title: "foo"}, nil
}

func (s *KobService) CreateEntry(context.Context, *rpc.CreateEntryRequest) (*rpc.Entry, error) {
	return &rpc.Entry{Title: "foo"}, nil
}

func (s *KobService) UpdateEntry(context.Context, *rpc.UpdateEntryRequest) (*rpc.Entry, error) {
	return &rpc.Entry{Title: "foo"}, nil
}

func (s *KobService) ListEntryComments(context.Context, *rpc.ListEntryCommentsRequest) (*rpc.CommentList, error) {
	return &rpc.CommentList{Comments: []*rpc.Comment{{Content: "hello"}}}, nil
}

func (s *KobService) CreateEntryComment(context.Context, *rpc.CreateEntryCommentRequest) (*rpc.Comment, error) {
	return &rpc.Comment{Content: "ok"}, nil
}

func (s *KobService) UpdateEntryComment(context.Context, *rpc.UpdateEntryCommentRequest) (*rpc.Comment, error) {
	return &rpc.Comment{Content: "ok"}, nil
}

func (s *KobService) CreateJWT(context.Context, *empty.Empty) (*rpc.JWT, error) {
	//https://skarlso.github.io/2016/06/12/google-signin-with-go/
	return nil, nil
}

func (s *KobService) CreateAuthURL(context.Context, *empty.Empty) (*rpc.AuthURL, error) {
	//https://skarlso.github.io/2016/06/12/google-signin-with-go/
	return nil, nil
}
