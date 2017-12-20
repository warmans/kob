package service

import "context"
import "github.com/warmans/kob/server/rpc"

func NewKobService() *KobService {
	return &KobService{}
}

type KobService struct{}

func (s *KobService) GetEntry(context.Context, *rpc.GetEntryRequest) (*rpc.Entry, error) {
	return &rpc.Entry{Title: "foo"}, nil
}

func (s *KobService) ListEntries(context.Context, *rpc.ListEntriesRequest) (*rpc.EntryList, error) {
	return &rpc.EntryList{Entries: []*rpc.Entry{{Title: "bar"}}}, nil
}
