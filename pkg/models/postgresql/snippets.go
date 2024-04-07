package postgresql

import (
	"alexedwards.net/snippetbox/pkg/models"
	"github.com/jmoiron/sqlx"
)

type SnippetModel struct {
	DB *sqlx.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
