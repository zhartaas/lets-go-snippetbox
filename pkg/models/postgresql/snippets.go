package postgresql

import (
	"alexedwards.net/snippetbox/pkg/models"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
)

type SnippetModel struct {
	DB *sqlx.DB
}

func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	expiration := time.Now().Add(time.Duration(expires) * 24 * time.Hour)
	stmt := `INSERT INTO snippets (title, content, expires) VALUES($1, $2, $3) RETURNING id`

	var id int
	err := m.DB.QueryRow(stmt, title, content, expiration).Scan(&id)

	if err != nil {
		return 0, nil
	}
	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > CURRENT_TIMESTAMP AND id=$1`

	// Use the QueryRow() method on the connection pool to execute our
	// SQL statement, passing in the untrusted id variable as the value for the
	// placeholder parameter. This returns a pointer to a sql.Row object which
	// holds the result from the database.
	row := m.DB.QueryRow(stmt, id)

	s := &models.Snippet{}

	// row.Scan() to copy the values from each field in sql.Row to the Snippet struct
	//Notice that the arguments to row.Scan are *pointers*
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err != nil {
		// If the query returns no rows, then row.Scan() will return a
		// sql.ErrNoRows error
		//We use the errors.Is() function check for that
		// error specifically, and return our own models.ErrNoRecord error
		// instead.
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
