package repositories

import (
	"github.com/jai-rathore/lists/logger"
	"github.com/jai-rathore/lists/models"
	"github.com/jmoiron/sqlx"
)

type IListRepository interface {
	GetListByID(id int) (*models.List, error)
	CreateList(list *models.List) (*models.List, error)
}

// ListRepository handles list requests
type ListRepository struct {
	DbContext *sqlx.DB
}

// NewListRepository creates a new list repository
func NewListRepository() *ListRepository {
	return &ListRepository{DbContext: db}
}

// CreateList creates a new list
func (r *ListRepository) CreateList(list *models.List) (*models.List, error) {
	// insert the list, if insert is successful, get the list by latest insert id and return it
	// otherwise return error
	query := `INSERT INTO lists (title, description, userid) VALUES (?, ?, ?)`
	result, err := r.DbContext.Exec(query, list.Title, list.Description, list.UserID)
	if err != nil {
		logger.Error("failed to create list " + err.Error())
		return nil, err
	}
	id, err := result.LastInsertId()
	//convert int64 to int
	list.ID = int(id)
	if err != nil {
		logger.Error("failed to get last insert id " + err.Error())
		return nil, err
	}
	return list, nil
}

// GetListByID returns a list by id
func (r *ListRepository) GetListByID(id int) (*models.List, error) {
	query := `SELECT id, title, description, userid, createdat, updatedat FROM lists WHERE id=?`
	list := &models.List{}
	err := r.DbContext.Get(list, query, id)
	if err != nil {
		logger.Error("failed to get list by id " + err.Error())
		return nil, err
	}
	return list, nil
}
