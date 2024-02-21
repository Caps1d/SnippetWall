package models

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
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
	var s Snippet
	query := `
            SELECT id, title, content, created, expires 
            FROM snippets
            WHERE id = $1 AND expires > CURRENT_TIMESTAMP;`

	err := m.DB.QueryRow(context.Background(), query, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if errors.Is(err, sql.ErrNoRows) || s.ID != id {
		// return our custom error
		return nil, ErrNoRecord
	}

	return &s, nil
}

// will return 10 most recent snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {

	query := `
            SELECT * 
            FROM snippets
            WHERE expires > CURRENT_TIMESTAMP
            ORDER BY id DESC LIMIT 10;
  `

	rows, err := m.DB.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snippets := []*Snippet{}

	for rows.Next() {
		s := &Snippet{}

		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
