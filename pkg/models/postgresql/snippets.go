package postgresql

import (
	"alexedwards.net/snippetbox/pkg/models"
	"github.com/jmoiron/sqlx"
	"time"
)

type SnippetModel struct {
	DB *sqlx.DB
}

func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {

	// SQL statement we want to execute. Important note in postgresql placeholder looks like "$1", "$2" and etc
	// instead of "?" in mySql

	expiration := time.Now().Add(time.Duration(expires) * 24 * time.Hour)
	stmt := `INSERT INTO snippets (title, content, expires) VALUES($1, $2, $3) RETURNING id`

	// Exec method to execute the statement
	// first param is sql statement followed by variables
	var id int
	err := m.DB.QueryRow(stmt, title, content, expiration).Scan(&id)

	if err != nil {
		return 0, nil
	}
	// LastInsertID() get the id of last inserted snippet
	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
