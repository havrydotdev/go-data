package tests

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/gavrylenkoIvan/go-data"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	tc "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var r *data.SQLRepository[int, Test]

var pgx *tc.PostgresContainer

type Test struct {
	ID    int `db:"id"`
	Name  string
	Count int
}

func (t Test) GetTableName() string {
	return "test"
}

func setup() {
	ctx := context.Background()
	pq, err := tc.RunContainer(ctx,
		testcontainers.WithImage("postgres:latest"),
		tc.WithDatabase("postgres"),
		tc.WithUsername("user"),
		tc.WithPassword("password"),
		tc.WithInitScripts(filepath.Join("schema", "up.sh")),
		testcontainers.WithWaitStrategy(wait.ForLog("database system is ready to accept connections").WithOccurrence(2)),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = pq.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
	url, err := pq.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}

	repo := data.NewSQLRepository[int, Test](db)
	r = repo
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestGetAll(t *testing.T) {
	res, err := r.GetAll()
	if err != nil {
		t.Error(err)
	}

	if res == nil {
		t.Error("res is nil")
	}
	js, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}
	println(string(js))
}

func TestGetById(t *testing.T) {
	res, err := r.GetById(1)
	if err != nil {
		t.Error(err)
	}

	js, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}
	println(string(js))
}
