package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Snippet struct {
	ID          int
	Title       string
	Content     string
	CreatedTime time.Time
	Expires     time.Time // this is for GET requests
	ExpiresInt  int       // this is for inserts POST requests
}

type SnippetModel struct {
	DB *pgxpool.Pool
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int64, error) {
	query := `
            INSERT INTO snippets (title, content, created, expires)
            VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '1' day * $3)
            RETURNING id;`

	var id int64
	err := m.DB.QueryRow(context.Background(), query, title, content, expires).Scan(&id)
	return id, err
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	// snippet := new snippet
	return nil, nil
}

// will return 10 most recent snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
