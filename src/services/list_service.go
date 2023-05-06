package services

import (
	"github.com/jai-rathore/lists/logger"
	"github.com/jai-rathore/lists/models"
	"github.com/jai-rathore/lists/repositories"
)

type IListService interface {
	GetListById(id int) (*models.List, error)
	CreateList(list *models.List) (*models.List, error)
}

// ListService handles list requests
type ListService struct {
	listRepository repositories.IListRepository
}

// NewListService creates a new list service
func NewListService(listRepository repositories.IListRepository) *ListService {
	return &ListService{listRepository: listRepository}
}

// GetListById returns a list by id
func (l *ListService) GetListById(id int) (*models.List, error) {
	list, err := l.listRepository.GetListByID(id)
	if err != nil {
		logger.Error("failed to get list by id " + err.Error())
		return nil, err
	}
	return list, nil
}

// CreateList creates a new list
func (l *ListService) CreateList(list *models.List) (*models.List, error) {
	listResp, err := l.listRepository.CreateList(list)
	if err != nil {
		logger.Error("failed to create list " + err.Error())
		return nil, err
	}
	return listResp, nil
}
