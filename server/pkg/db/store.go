package db

import "github.com/gocraft/dbr"
import "github.com/warmans/kob/server/pkg/rpc"
import "database/sql"
import "github.com/pkg/errors"
import sq "github.com/Masterminds/squirrel"

type Scannable interface {
	Scan(dest ...interface{}) error
}

func NewStore(conn *dbr.Connection) *Store {
	return &Store{conn: conn}
}

type Store struct {
	conn *dbr.Connection
}

func (s *Store) Session() (*Session, error) {
	tx, err := s.conn.NewSession(nil).Begin()
	return &Session{tx: tx}, err
}

func (s *Store) WithSession(f func(*Session) error) error {
	sess, err := s.Session()
	if err != nil {
		return err
	}
	if err := f(sess); err != nil {
		if err2 := sess.Rollback(); err2 != nil {
			return errors.Wrapf(err, "with rollback failure: %s")
		}
		return err
	}
	return sess.Commit()
}

type Session struct {
	tx *dbr.Tx
}

func (s *Session) Commit() error {
	return s.tx.Commit()
}

func (s *Session) Rollback() error {
	return s.tx.Rollback()
}

func (s *Session) CreateEntry(entry *rpc.CreateEntryRequest) (*rpc.Entry, error) {
	var id int64
	if err := s.tx.QueryRow("INSERT INTO entry (title, content) VALUES (?, ?) RETURNING (id)", entry.Title, entry.Content).Scan(&id); err != nil {
		return nil, err
	}
	return s.GetEntry(id)
}

func (s *Session) GetEntry(id int64) (*rpc.Entry, error) {
	entries, err := s.entries(1, 0, sq.Expr("e.id = ?", id))
	if err != nil {
		return nil, err
	}
	if len(entries.Entries) == 0 {
		return nil, sql.ErrNoRows
	}
	return entries.Entries[0], nil
}

func (s *Session) ListEntries(req *rpc.ListEntriesRequest) (*rpc.EntryList, error) {
	return s.entries(uint64(req.NumPerPage), uint64(req.NumPerPage*req.Page))
}

func (s *Session) entries(limit uint64, offset uint64, where ...sq.Sqlizer) (*rpc.EntryList, error) {
	q := sq.Select("e.id", "e.title", "e.content").From("entry e").Limit(limit).Offset(offset)
	for _, c := range where {
		q.Where(c)
	}
	sql, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	res, err := s.tx.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	entries := &rpc.EntryList{Entries: make([]*rpc.Entry, 0)}

	if entries.NumResults, err = s.entryCount(where...); err != nil {
		return nil, err
	}
	if entries.NumResults == 0 {
		return entries, nil //if the count is 0 skip second query
	}
	for res.Next() {
		entry := &rpc.Entry{}
		if err := res.Scan(entry.Id, entry.Title, entry.Content); err != nil {
			return nil, err
		}
		entries.Entries = append(entries.Entries, entry)
	}
	return entries, nil
}

func (s *Session) entryCount(where ...sq.Sqlizer) (count int64, err error) {
	q := sq.Select("COUNT(*)").From("entry e")
	for _, c := range where {
		q.Where(c)
	}
	sql, args, err := q.ToSql()
	if err != nil {
		return
	}
	err = s.tx.QueryRow(sql, args...).Scan(&count)
	if err != nil {
		return
	}
	return
}
