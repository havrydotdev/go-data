# Golang Data

### This project was inspired by [Spring Data](https://spring.io/projects/spring-data)

## Quick start

**You can use data.SQLRepository instead of your custom repository:**

```go
import (
    _ "github.com/lib/pq"
    "github.com/gavrylenkoIvan/go-data"
)

const (
    driver = "postgres"
    url = "postgres_url"
)

// Define struct that will be used in SQLRepository
type Note struct {
	ID    int
	Text  string
}

// This method is required to use SQLRepository
func (n Note) GetTableName() string {
	return "test"
}

func main() {
    // Open db connection
    db, err := sql.Open(driver, url)
	if err != nil {
		log.Fatal(err)
	}

    // Define SQLRepository
    repo := data.NewSQLRepository[int, Note](db)
    
    // You are ready to go!
    res, err := r.GetById(1) 
}
```

**Or you can use it as wrapper for your custom repository:**

```go
import (
    _ "github.com/lib/pq"
    "github.com/gavrylenkoIvan/go-data"
)

const (
    driver = "postgres"
    url = "postgres_url"
)

// Create struct that will be used in SQLRepository
type Note struct {
	ID    int
	Text  string
}

// This method is required to use SQLRepository
func (n Note) GetTableName() string {
	return "test"
}

// Define custom repo
type NoteRepo struct {
    db *sqlx.DB
}

func NewNoteRepo(db *sqlx.DB) *NoteRepo {
    return &NoteRepo{
        db: db,
    }
}

func (r *NoteRepo) GetByText() (Note, err) {
    // Implement method
}

func main() {
    // Open db connection
    db, err := sql.Open(driver, url)
	if err != nil {
		log.Fatal(err)
	}

    // Init custom repository
    customRepo := NewNoteRepo(db)

    // Init SQLRepository with NoteRepo as child
    repo := data.NewSQLRepositoryWithChild[int, Note](customRepo, db)
}
```