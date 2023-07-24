package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllQuery(t *testing.T) {
	testTable := []struct {
		Table          string
		ExpectedResult string
	}{
		{
			Table:          "test",
			ExpectedResult: "SELECT * FROM test",
		},
		{
			Table:          "32i48jkd,mncv9dn3mer",
			ExpectedResult: "SELECT * FROM 32i48jkd,mncv9dn3mer",
		},
		{
			Table:          "192043204239",
			ExpectedResult: "SELECT * FROM 192043204239",
		},
		{
			Table:          "notes",
			ExpectedResult: "SELECT * FROM notes",
		},
	}

	for _, tt := range testTable {
		assert.Equal(t, tt.ExpectedResult, getAll(tt.Table))
	}
}

func TestGetByIdQuery(t *testing.T) {
	testTable := []struct {
		Table          string
		ExpectedResult string
	}{
		{
			Table:          "test",
			ExpectedResult: "SELECT * FROM test WHERE id = $1",
		},
		{
			Table:          "32i48jkd,mncv9dn3mer",
			ExpectedResult: "SELECT * FROM 32i48jkd,mncv9dn3mer WHERE id = $1",
		},
		{
			Table:          "192043204239",
			ExpectedResult: "SELECT * FROM 192043204239 WHERE id = $1",
		},
		{
			Table:          "notes",
			ExpectedResult: "SELECT * FROM notes WHERE id = $1",
		},
	}

	for _, tt := range testTable {
		assert.Equal(t, tt.ExpectedResult, getById(tt.Table))
	}
}

func TestInsertQuery(t *testing.T) {
	testTable := []struct {
		Table          string
		Columns        []string
		ExpectedResult string
	}{
		{
			Table:          "test",
			Columns:        []string{"text", "value"},
			ExpectedResult: "INSERT INTO test (text, value) VALUES ($1, $2) RETURNING id",
		},
		{
			Table:          "hello",
			Columns:        []string{"123", "333", "hello", "simple"},
			ExpectedResult: "INSERT INTO hello (123, 333, hello, simple) VALUES ($1, $2, $3, $4) RETURNING id",
		},
	}

	for _, tt := range testTable {
		assert.Equal(t, tt.ExpectedResult, insert(tt.Table, tt.Columns))
	}
}
