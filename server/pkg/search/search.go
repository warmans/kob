package search

import "github.com/blevesearch/bleve"
import bs "github.com/blevesearch/bleve/search"
import "github.com/warmans/kob/server/pkg/rpc"
import "fmt"
import "bytes"

func NewIndex(path string) (*Index, error) {

	index, err := bleve.Open(path)
	if err != nil {
		if err == bleve.ErrorIndexPathDoesNotExist {
			index, err = createIndex(path)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return &Index{index: index}, nil
}

func createIndex(path string) (bleve.Index, error) {
	index, err := bleve.New(path, bleve.NewIndexMapping())
	if err != nil {
		return nil, err
	}
	return index, err
}

type Index struct {
	index bleve.Index
}

func (i *Index) IndexEntry(entry *rpc.Entry) error {
	return i.index.Index(fmt.Sprintf("%d", entry.Id), entry)
}

func (i *Index) Search(req *bleve.SearchRequest) (*bleve.SearchResult, error) {
	return i.index.Search(req)
}

func MakePreview(result *bs.DocumentMatch) string {
	previewBuff := bytes.NewBufferString("")
	for _, fragments := range result.Fragments {
		for _, fragment := range fragments {
			fmt.Fprintf(previewBuff, "%s", fragment)
		}
	}
	for otherFieldName, otherFieldValue := range result.Fields {
		if _, ok := result.Fragments[otherFieldName]; !ok {
			fmt.Fprintf(previewBuff, "%v", otherFieldValue)
		}
	}
	return previewBuff.String()
}