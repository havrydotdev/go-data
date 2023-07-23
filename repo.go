package data

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DataTable interface {
	GetTableName() string
}

type SQLRepository[K comparable, T DataTable] struct {
	repo any
	db   *sqlx.DB
}

func NewSQLRepository[K comparable, T DataTable](db *sql.DB) *SQLRepository[K, T] {
	return &SQLRepository[K, T]{
		db: sqlx.NewDb(db, "postgres"),
	}
}

func NewSQLRepositoryWithChild[K comparable, T DataTable](customRepository any, db *sql.DB) *SQLRepository[K, T] {
	return &SQLRepository[K, T]{
		repo: customRepository,
		db:   sqlx.NewDb(db, "postgres"),
	}
}
