// IMPORTANT! This is auto generated code by https://github.com/src-d/go-kallax
// Please, do not touch the code below, and if you do, do it under your own
// risk. Take into account that all the code you write here will be completely
// erased from earth the next time you generate the kallax models.
package model

import (
	"database/sql"
	"fmt"

	"github.com/src-d/go-kallax"
	"github.com/src-d/go-kallax/types"
)

var _ types.SQLType
var _ fmt.Formatter

// NewMention returns a new instance of Mention.
func NewMention() (record *Mention) {
	record = &Mention{}
	if record != nil {
		record.SetID(kallax.NewID())
	}
	return
}

func (r *Mention) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return &r.Model.ID, nil
	case "created_at":
		return &r.Timestamps.CreatedAt, nil
	case "updated_at":
		return &r.Timestamps.UpdatedAt, nil
	case "endpoint":
		return &r.Endpoint, nil
	case "provider":
		return &r.Provider, nil
	case "vcs":
		return &r.VCS, nil
	case "context":
		return types.JSON(&r.Context), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Mention: %s", col)
	}
}

func (r *Mention) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.Model.ID, nil
	case "created_at":
		return r.Timestamps.CreatedAt, nil
	case "updated_at":
		return r.Timestamps.UpdatedAt, nil
	case "endpoint":
		return r.Endpoint, nil
	case "provider":
		return r.Provider, nil
	case "vcs":
		return (string)(r.VCS), nil
	case "context":
		return types.JSON(r.Context), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Mention: %s", col)
	}
}

func (r *Mention) NewRelationshipRecord(field string) (kallax.Record, error) {
	return nil, fmt.Errorf("kallax: model Mention has no relationships")
}

func (r *Mention) SetRelationship(field string, rel interface{}) error {
	return fmt.Errorf("kallax: model Mention has no relationships")
}

// MentionStore is the entity to access the records of the type Mention
// in the database.
type MentionStore struct {
	*kallax.Store
}

// NewMentionStore creates a new instance of MentionStore
// using a SQL database.
func NewMentionStore(db *sql.DB) *MentionStore {
	return &MentionStore{kallax.NewStore(db)}
}

// Insert inserts a Mention in the database. A non-persisted object is
// required for this operation.
func (s *MentionStore) Insert(record *Mention) error {

	if err := record.BeforeSave(); err != nil {
		return err
	}

	return s.Store.Insert(Schema.Mention.BaseSchema, record)

}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *MentionStore) Update(record *Mention, cols ...kallax.SchemaField) (updated int64, err error) {

	if err := record.BeforeSave(); err != nil {
		return 0, err
	}

	return s.Store.Update(Schema.Mention.BaseSchema, record, cols...)

}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *MentionStore) Save(record *Mention) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *MentionStore) Delete(record *Mention) error {

	return s.Store.Delete(Schema.Mention.BaseSchema, record)

}

// Find returns the set of results for the given query.
func (s *MentionStore) Find(q *MentionQuery) (*MentionResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewMentionResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *MentionStore) MustFind(q *MentionQuery) *MentionResultSet {
	return NewMentionResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *MentionStore) Count(q *MentionQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *MentionStore) MustCount(q *MentionQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *MentionStore) FindOne(q *MentionQuery) (*Mention, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *MentionStore) MustFindOne(q *MentionQuery) *Mention {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Mention with the data in the database and
// makes it writable.
func (s *MentionStore) Reload(record *Mention) error {
	return s.Store.Reload(Schema.Mention.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *MentionStore) Transaction(callback func(*MentionStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&MentionStore{store})
	})
}

// MentionQuery is the object used to create queries for the Mention
// entity.
type MentionQuery struct {
	*kallax.BaseQuery
}

// NewMentionQuery returns a new instance of MentionQuery.
func NewMentionQuery() *MentionQuery {
	return &MentionQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Mention.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *MentionQuery) Select(columns ...kallax.SchemaField) *MentionQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *MentionQuery) SelectNot(columns ...kallax.SchemaField) *MentionQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *MentionQuery) Copy() *MentionQuery {
	return &MentionQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *MentionQuery) Order(cols ...kallax.ColumnOrder) *MentionQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *MentionQuery) BatchSize(size uint64) *MentionQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *MentionQuery) Limit(n uint64) *MentionQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *MentionQuery) Offset(n uint64) *MentionQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *MentionQuery) Where(cond kallax.Condition) *MentionQuery {
	q.BaseQuery.Where(cond)
	return q
}

// MentionResultSet is the set of results returned by a query to the
// database.
type MentionResultSet struct {
	ResultSet kallax.ResultSet
	last      *Mention
	lastErr   error
}

// NewMentionResultSet creates a new result set for rows of the type
// Mention.
func NewMentionResultSet(rs kallax.ResultSet) *MentionResultSet {
	return &MentionResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *MentionResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Mention.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Mention)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Mention")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *MentionResultSet) Get() (*Mention, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *MentionResultSet) ForEach(fn func(*Mention) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *MentionResultSet) All() ([]*Mention, error) {
	var result []*Mention
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *MentionResultSet) One() (*Mention, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *MentionResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *MentionResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewRepository returns a new instance of Repository.
func NewRepository() (record *Repository) {
	record = newRepository()
	if record != nil {
		record.SetID(kallax.NewID())
	}
	return
}

func (r *Repository) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return &r.Model.ID, nil
	case "created_at":
		return &r.Timestamps.CreatedAt, nil
	case "updated_at":
		return &r.Timestamps.UpdatedAt, nil
	case "endpoints":
		return types.Slice(&r.Endpoints), nil
	case "status":
		return &r.Status, nil
	case "fetched_at":
		return r.FetchedAt, nil
	case "fetch_error_at":
		return r.FetchErrorAt, nil
	case "last_commit_at":
		return r.LastCommitAt, nil
	case "_references":
		return types.JSON(&r.References), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Repository: %s", col)
	}
}

func (r *Repository) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.Model.ID, nil
	case "created_at":
		return r.Timestamps.CreatedAt, nil
	case "updated_at":
		return r.Timestamps.UpdatedAt, nil
	case "endpoints":
		return types.Slice(r.Endpoints), nil
	case "status":
		return (string)(r.Status), nil
	case "fetched_at":
		return r.FetchedAt, nil
	case "fetch_error_at":
		return r.FetchErrorAt, nil
	case "last_commit_at":
		return r.LastCommitAt, nil
	case "_references":
		return types.JSON(r.References), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Repository: %s", col)
	}
}

func (r *Repository) NewRelationshipRecord(field string) (kallax.Record, error) {
	return nil, fmt.Errorf("kallax: model Repository has no relationships")
}

func (r *Repository) SetRelationship(field string, rel interface{}) error {
	return fmt.Errorf("kallax: model Repository has no relationships")
}

// RepositoryStore is the entity to access the records of the type Repository
// in the database.
type RepositoryStore struct {
	*kallax.Store
}

// NewRepositoryStore creates a new instance of RepositoryStore
// using a SQL database.
func NewRepositoryStore(db *sql.DB) *RepositoryStore {
	return &RepositoryStore{kallax.NewStore(db)}
}

// Insert inserts a Repository in the database. A non-persisted object is
// required for this operation.
func (s *RepositoryStore) Insert(record *Repository) error {

	if err := record.BeforeSave(); err != nil {
		return err
	}

	return s.Store.Insert(Schema.Repository.BaseSchema, record)

}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *RepositoryStore) Update(record *Repository, cols ...kallax.SchemaField) (updated int64, err error) {

	if err := record.BeforeSave(); err != nil {
		return 0, err
	}

	return s.Store.Update(Schema.Repository.BaseSchema, record, cols...)

}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *RepositoryStore) Save(record *Repository) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *RepositoryStore) Delete(record *Repository) error {

	return s.Store.Delete(Schema.Repository.BaseSchema, record)

}

// Find returns the set of results for the given query.
func (s *RepositoryStore) Find(q *RepositoryQuery) (*RepositoryResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewRepositoryResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *RepositoryStore) MustFind(q *RepositoryQuery) *RepositoryResultSet {
	return NewRepositoryResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *RepositoryStore) Count(q *RepositoryQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *RepositoryStore) MustCount(q *RepositoryQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *RepositoryStore) FindOne(q *RepositoryQuery) (*Repository, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *RepositoryStore) MustFindOne(q *RepositoryQuery) *Repository {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Repository with the data in the database and
// makes it writable.
func (s *RepositoryStore) Reload(record *Repository) error {
	return s.Store.Reload(Schema.Repository.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *RepositoryStore) Transaction(callback func(*RepositoryStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&RepositoryStore{store})
	})
}

// RepositoryQuery is the object used to create queries for the Repository
// entity.
type RepositoryQuery struct {
	*kallax.BaseQuery
}

// NewRepositoryQuery returns a new instance of RepositoryQuery.
func NewRepositoryQuery() *RepositoryQuery {
	return &RepositoryQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Repository.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *RepositoryQuery) Select(columns ...kallax.SchemaField) *RepositoryQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *RepositoryQuery) SelectNot(columns ...kallax.SchemaField) *RepositoryQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *RepositoryQuery) Copy() *RepositoryQuery {
	return &RepositoryQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *RepositoryQuery) Order(cols ...kallax.ColumnOrder) *RepositoryQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *RepositoryQuery) BatchSize(size uint64) *RepositoryQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *RepositoryQuery) Limit(n uint64) *RepositoryQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *RepositoryQuery) Offset(n uint64) *RepositoryQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *RepositoryQuery) Where(cond kallax.Condition) *RepositoryQuery {
	q.BaseQuery.Where(cond)
	return q
}

// RepositoryResultSet is the set of results returned by a query to the
// database.
type RepositoryResultSet struct {
	ResultSet kallax.ResultSet
	last      *Repository
	lastErr   error
}

// NewRepositoryResultSet creates a new result set for rows of the type
// Repository.
func NewRepositoryResultSet(rs kallax.ResultSet) *RepositoryResultSet {
	return &RepositoryResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *RepositoryResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Repository.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Repository)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Repository")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *RepositoryResultSet) Get() (*Repository, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *RepositoryResultSet) ForEach(fn func(*Repository) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *RepositoryResultSet) All() ([]*Repository, error) {
	var result []*Repository
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *RepositoryResultSet) One() (*Repository, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *RepositoryResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *RepositoryResultSet) Close() error {
	return rs.ResultSet.Close()
}

type schema struct {
	Mention    *schemaMention
	Repository *schemaRepository
}

type schemaMention struct {
	*kallax.BaseSchema
	ID        kallax.SchemaField
	CreatedAt kallax.SchemaField
	UpdatedAt kallax.SchemaField
	Endpoint  kallax.SchemaField
	Provider  kallax.SchemaField
	VCS       kallax.SchemaField
	Context   kallax.SchemaField
}

type schemaRepository struct {
	*kallax.BaseSchema
	ID           kallax.SchemaField
	CreatedAt    kallax.SchemaField
	UpdatedAt    kallax.SchemaField
	Endpoints    kallax.SchemaField
	Status       kallax.SchemaField
	FetchedAt    kallax.SchemaField
	FetchErrorAt kallax.SchemaField
	LastCommitAt kallax.SchemaField
	References   *schemaRepositoryReferences
}

type schemaRepositoryReferences struct {
	*kallax.BaseSchemaField
	CreatedAt kallax.SchemaField
	UpdatedAt kallax.SchemaField
	Name      kallax.SchemaField
	Hash      kallax.SchemaField
	Init      kallax.SchemaField
	Roots     kallax.SchemaField
}

func (s *schemaRepositoryReferences) At(n int) *schemaRepositoryReferences {
	return &schemaRepositoryReferences{
		BaseSchemaField: kallax.NewSchemaField("_references").(*kallax.BaseSchemaField),
		CreatedAt:       kallax.NewJSONSchemaKey(kallax.JSONAny, "_references", "CreatedAt"),
		UpdatedAt:       kallax.NewJSONSchemaKey(kallax.JSONAny, "_references", "UpdatedAt"),
		Name:            kallax.NewJSONSchemaKey(kallax.JSONText, "_references", fmt.Sprint(n), "Name"),
		Hash:            kallax.NewJSONSchemaArray("_references", fmt.Sprint(n), "Hash"),
		Init:            kallax.NewJSONSchemaArray("_references", fmt.Sprint(n), "Init"),
		Roots:           kallax.NewJSONSchemaArray("_references", fmt.Sprint(n), "Roots"),
	}
}

type schemaRepositoryReferencesTimestamps struct {
	*kallax.BaseSchemaField
	CreatedAt kallax.SchemaField
	UpdatedAt kallax.SchemaField
}

var Schema = &schema{
	Mention: &schemaMention{
		BaseSchema: kallax.NewBaseSchema(
			"mentions",
			"__mention",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			func() kallax.Record {
				return new(Mention)
			},
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("created_at"),
			kallax.NewSchemaField("updated_at"),
			kallax.NewSchemaField("endpoint"),
			kallax.NewSchemaField("provider"),
			kallax.NewSchemaField("vcs"),
			kallax.NewSchemaField("context"),
		),
		ID:        kallax.NewSchemaField("id"),
		CreatedAt: kallax.NewSchemaField("created_at"),
		UpdatedAt: kallax.NewSchemaField("updated_at"),
		Endpoint:  kallax.NewSchemaField("endpoint"),
		Provider:  kallax.NewSchemaField("provider"),
		VCS:       kallax.NewSchemaField("vcs"),
		Context:   kallax.NewSchemaField("context"),
	},
	Repository: &schemaRepository{
		BaseSchema: kallax.NewBaseSchema(
			"repositories",
			"__repository",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			func() kallax.Record {
				return new(Repository)
			},
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("created_at"),
			kallax.NewSchemaField("updated_at"),
			kallax.NewSchemaField("endpoints"),
			kallax.NewSchemaField("status"),
			kallax.NewSchemaField("fetched_at"),
			kallax.NewSchemaField("fetch_error_at"),
			kallax.NewSchemaField("last_commit_at"),
			kallax.NewSchemaField("_references"),
		),
		ID:           kallax.NewSchemaField("id"),
		CreatedAt:    kallax.NewSchemaField("created_at"),
		UpdatedAt:    kallax.NewSchemaField("updated_at"),
		Endpoints:    kallax.NewSchemaField("endpoints"),
		Status:       kallax.NewSchemaField("status"),
		FetchedAt:    kallax.NewSchemaField("fetched_at"),
		FetchErrorAt: kallax.NewSchemaField("fetch_error_at"),
		LastCommitAt: kallax.NewSchemaField("last_commit_at"),
		References: &schemaRepositoryReferences{
			BaseSchemaField: kallax.NewSchemaField("_references").(*kallax.BaseSchemaField),
			CreatedAt:       kallax.NewJSONSchemaKey(kallax.JSONAny, "_references", "CreatedAt"),
			UpdatedAt:       kallax.NewJSONSchemaKey(kallax.JSONAny, "_references", "UpdatedAt"),
			Name:            kallax.NewJSONSchemaKey(kallax.JSONText, "_references", "Name"),
			Hash:            kallax.NewJSONSchemaArray("_references", "Hash"),
			Init:            kallax.NewJSONSchemaArray("_references", "Init"),
			Roots:           kallax.NewJSONSchemaArray("_references", "Roots"),
		},
	},
}
