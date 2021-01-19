package mysql

import (
	"database/sql"
	"trannguyenhung011086/learn-go-web/pkg/models"
)

// UserModel : database model for User
type UserModel struct {
	DB *sql.DB
}

// Insert - insert new user
func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// Authenticate - authenticate user and return user id
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get - get user by id
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
