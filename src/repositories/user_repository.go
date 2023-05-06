package repositories

import (
	"database/sql"
	"fmt"

	"github.com/jai-rathore/lists/models"
	"github.com/jmoiron/sqlx"
)

type IUserRepository interface {
	GetUserByID(id int) (*models.User, error)
}

type UserRepository struct {
	DbContext *sqlx.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	query := `SELECT id, email, created_at, updated_at FROM users WHERE id=?`
	user := &models.User{}
	err := r.DbContext.Get(user, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return nil, err
	}
	return user, nil
}
