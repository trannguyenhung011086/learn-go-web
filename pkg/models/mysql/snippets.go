package mysql

import (
	"database/sql"
	"trannguyenhung011086/learn-go-web/pkg/models"
)

// SnippetModel : database model for snippet
type SnippetModel struct {
	DB *sql.DB
}

// Insert : insert snippet
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Get : get snippet by id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Latest : get number of latest snippets
func (m *SnippetModel) Latest(num int) ([]*models.Snippet, error) {
	return nil, nil
}
