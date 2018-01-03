package service

import "context"
import "github.com/warmans/kob/server/pkg/rpc"
import "github.com/golang/protobuf/ptypes/empty"
import "github.com/warmans/kob/server/pkg/db"

func NewKobService(store *db.Store) *KobService {
	return &KobService{store: store}
}

type KobService struct {
	store *db.Store
}

func (s *KobService) ListEntries(ctx context.Context, req *rpc.ListEntriesRequest) (entries *rpc.EntryList, err error) {
	err = s.store.WithSession(func(s *db.Session) error {
		entries, err = s.ListEntries(req)
		return err
	})
	return entries, Err2Code(err)
}

func (s *KobService) GetEntry(ctx context.Context, req *rpc.GetEntryRequest) (entry *rpc.Entry, err error) {
	err = s.store.WithSession(func(s *db.Session) error {
		entry, err = s.GetEntry(req.Id)
		return err
	})
	return entry, Err2Code(err)
}

func (s *KobService) CreateEntry(ctx context.Context, req *rpc.CreateEntryRequest) (entry *rpc.Entry, err error) {
	req.AuthorId = 1; //todo: from token
	err = s.store.WithSession(func(s *db.Session) error {
		entry, err = s.CreateEntry(req)
		return err
	})
	return entry, Err2Code(err)
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

func (s *KobService) CreateJWT(ctx context.Context, author *rpc.Author) (*rpc.JWT, error) {
	//https://skarlso.github.io/2016/06/12/google-signin-with-go/
	return nil, nil
}

func (s *KobService) CreateAuthURL(context.Context, *empty.Empty) (*rpc.AuthURL, error) {
	return nil, nil
}

func (s *KobService) ListActivity(context.Context, *empty.Empty) (*rpc.ActivityList, error) {
	return nil, nil
}
