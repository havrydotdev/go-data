package data

import (
	"database/sql"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
)

type DataTable interface {
	GetTableName() string
}

type Queries struct {
	getAll  string
	getById string
	insert  string
	// update  string
	// delete  string
}

type SQLRepository[K comparable, T DataTable] struct {
	repo    any
	db      *sqlx.DB
	qs      Queries
	table   string
	columns []string
}

func NewSQLRepository[K comparable, T DataTable](db *sql.DB) *SQLRepository[K, T] {
	r := &SQLRepository[K, T]{
		db:    sqlx.NewDb(db, "postgres"),
		table: T.GetTableName(*new(T)),
	}

	r.initQueries()

	return r
}

func NewSQLRepositoryWithChild[K comparable, T DataTable](customRepository any, db *sql.DB) *SQLRepository[K, T] {
	r := &SQLRepository[K, T]{
		repo:  customRepository,
		db:    sqlx.NewDb(db, "postgres"),
		table: T.GetTableName(*new(T)),
	}

	r.initQueries()

	return r
}

func (r *SQLRepository[K, T]) initQueries() {
	r.qs.getAll = getAll(r.table)
	r.qs.getById = getById(r.table)
	t := reflect.TypeOf(*new(T))
	keys := []string{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if strings.ToLower(field.Name) == "id" {
			continue
		}
		keys = append(keys, field.Name)
	}
	r.qs.insert = insert(r.table, keys)
	r.columns = keys
}
