package db

import "github.com/gocraft/dbr"
import "github.com/warmans/kob/server/pkg/rpc"
import "database/sql"
import "github.com/pkg/errors"
import sq "github.com/Masterminds/squirrel"
import "time"
import "github.com/warmans/kob/server/pkg/search"

type Scannable interface {
	Scan(dest ...interface{}) error
}

func NewStore(conn *dbr.Connection, index *search.Index) *Store {
	return &Store{conn: conn, index: index}
}

type Store struct {
	conn  *dbr.Connection
	index *search.Index
}

func (s *Store) Session() (*Session, error) {
	tx, err := s.conn.NewSession(nil).Begin()
	return &Session{tx: tx, index: s.index}, err
}

func (s *Store) WithSession(f func(*Session) error) error {
	sess, err := s.Session()
	if err != nil {
		return err
	}
	if err := f(sess); err != nil {
		if err2 := sess.Rollback(); err2 != nil {
			return errors.Wrapf(err, "with rollback failure: %s", err2)
		}
		return err
	}
	return sess.Commit()
}

type Session struct {
	tx *dbr.Tx
	index *search.Index
}

func (s *Session) Commit() error {
	return s.tx.Commit()
}

func (s *Session) Rollback() error {
	return s.tx.Rollback()
}

func (s *Session) CreateEntry(entry *rpc.CreateEntryRequest) (*rpc.Entry, error) {
	var id int64
	if err := s.tx.QueryRow(
		"INSERT INTO entry (title, content, author_id, created_date, updated_date) VALUES ($1, $2, $3, $4, $5) RETURNING (id)",
		entry.Title,
		entry.Content,
		entry.AuthorId,
		time.Now().Format(time.RFC3339),
		time.Now().Format(time.RFC3339),
	).Scan(&id); err != nil {
		return nil, errors.Wrap(err, "failed to create entry")
	}
	for _, t := range entry.TagIds {
		if _, err := s.tx.Exec("INSERT INTO entry_tags (entry_id, tag_id) VALUES ($1, $2)", id, t); err != nil {
			return nil, errors.Wrap(err, "failed to link tag "+t)
		}
	}
	storedEntry, err := s.GetEntry(id)
	if err != nil {
		if err := s.index.IndexEntry(storedEntry); err != nil {
			return nil, errors.Wrap(err, "indexing failed")
		}
	}
	return storedEntry, err

}

func (s *Session) UpdateEntry(entry *rpc.UpdateEntryRequest) (*rpc.Entry, error) {
	if _, err := s.tx.Exec(
		"UPDATE entry SET title=$1, content=$2, updated_date=$3 WHERE id=$4",
		entry.Title,
		entry.Content,
		time.Now().Format(time.RFC3339),
		entry.Id,
	); err != nil {
		return nil, errors.Wrap(err, "failed to update entry")
	}
	if _, err := s.tx.Exec("DELETE FROM entry_tags WHERE entry_id=$1", entry.Id); err != nil {
		return nil, errors.Wrap(err, "failed to remove tags")
	}
	for _, t := range entry.TagIds {
		if _, err := s.tx.Exec("INSERT INTO entry_tags (entry_id, tag_id) VALUES ($1, $2)", entry.Id, t); err != nil {
			return nil, errors.Wrap(err, "failed to link tag "+t)
		}
	}
	updatedEntry, err := s.GetEntry(entry.Id)
	if err != nil {
		if err := s.index.IndexEntry(updatedEntry); err != nil {
			return nil, errors.Wrap(err, "indexing failed")
		}
	}
	return updatedEntry, err
}

func (s *Session) GetEntry(id int64) (*rpc.Entry, error) {
	entries, err := s.entries(1, 0, sq.Expr("e.id = ?", id))
	if err != nil {
		return nil, err
	}
	if len(entries) == 0 {
		return nil, sql.ErrNoRows
	}
	return entries[0], nil
}

func (s *Session) ListEntries(req *rpc.ListEntriesRequest) (*rpc.EntryList, error) {

	var err error
	entryList := &rpc.EntryList{}

	if req.NumPerPage == 0 {
		req.NumPerPage = 50
	}
	entryList.NumResults, err = s.entryCount( /*todo: where */ )
	if err != nil {
		return nil, err
	}
	if entryList.NumResults == 0 {
		return entryList, nil //if the count is 0 skip main query
	}

	entryList.Entries, err = s.entries(uint64(req.NumPerPage), uint64(req.NumPerPage*req.Page) /*todo: where */)
	if err != nil {
		return nil, err
	}
	return entryList, nil
}

func (s *Session) UpsertAuthor(author *rpc.Author) (*rpc.Author, error) {

	err := s.tx.QueryRow("SELECT id FROM author WHERE email=$1", author.Email).Scan(&author.Id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err := s.tx.QueryRow(`
		INSERT INTO author (sub, name, given_name, family_name, profile, picture, email, email_verified, gender) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (email) DO UPDATE SET sub=$1, name=$2, given_name=$3, family_name=$4, profile=$5, picture=$6, email=$7, email_verified=$8, gender=$9
		RETURNING id`,
		author.Sub,
		author.Name,
		author.GivenName,
		author.FamilyName,
		author.Profile,
		author.Picture,
		author.Email,
		author.EmailVerified,
		author.Gender,
	).Scan(&author.Id); err != nil {
		return nil, errors.Wrap(err, "failed to create entry")
	}
	return author, nil
}

func (s *Session) Reindex() error {
	pageNum := uint64(0)
	pageSize := uint64(25)
	for {
		page, err := s.entries(pageSize, pageSize*pageNum)
		if err != nil {
			return errors.Wrapf(err, "indexing failed requesting page %d", pageNum)
		}
		if len(page) == 0 {
			return nil
		}
		for _, entry := range page {
			if err := s.index.IndexEntry(entry); err != nil {
				return errors.Wrapf(err, "indexing failed at page %d entry %d", pageNum, entry.Id)
			}
		}
		pageNum++
	}
}

func (s *Session) entries(limit uint64, offset uint64, where ...sq.Sqlizer) ([]*rpc.Entry, error) {

	q := sq.Select(
		"e.id",
		"e.title",
		"e.content",
		"e.created_date",
		"e.updated_date",
		"a.id",
		"a.sub",
		"a.name",
		"a.given_name",
		"a.family_name",
		"a.profile",
		"a.picture",
		"a.email",
		"a.email_verified",
		"a.gender",
	).
		From("entry e").
		LeftJoin("author a ON e.author_id = a.id").
		Limit(limit).
		Offset(offset)

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

	entries := make([]*rpc.Entry, 0)
	for res.Next() {
		entry := &rpc.Entry{
			Author: &rpc.Author{},
		}
		if err := res.Scan(
			&entry.Id,
			&entry.Title,
			&entry.Content,
			&entry.CreatedDate,
			&entry.UpdatedDate,
			&entry.Author.Id,
			&entry.Author.Sub,
			&entry.Author.Name,
			&entry.Author.GivenName,
			&entry.Author.FamilyName,
			&entry.Author.Profile,
			&entry.Author.Picture,
			&entry.Author.Email,
			&entry.Author.EmailVerified,
			&entry.Author.Gender,
		); err != nil {
			res.Close()
			return nil, err
		}
		entries = append(entries, entry)
	}
	if err := res.Close(); err != nil {
		return nil, err
	}

	//finally add tags to entries
	for _, entry := range entries {
		if entry.Tags, err = s.tags(sq.Expr("et.entry_id=?", entry.Id)); err != nil {
			return nil, err
		}
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

func (s *Session) tags(where ...sq.Sqlizer) ([]*rpc.Tag, error) {
	q := sq.Select("t.id", "t.label").From("tag t").LeftJoin("entry_tag et ON t.id = et.tag_id")
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

	tags := make([]*rpc.Tag, 0)
	for res.Next() {
		tag := &rpc.Tag{}
		if err := res.Scan(&tag.Id, &tag.Label); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}
