package services

import (
	"golang-ture/internal/models"
	"golang-ture/internal/repositories"
)

type Service struct {
	TodoItem
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		TodoItem: NewTodoItemService(repos.TodoItem),
	}
}

type TodoItem interface {
	Create(item models.TodoItem) (string, error)
	GetAll() ([]models.TodoItem, error)
	GetById(itemId string) (models.TodoItem, error)
	Delete(itemId string) error
	Update(itemId string, input models.UpdateTodoItemInput) error
}
