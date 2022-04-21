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
	Create(item models.TodoItem) (int, error)
	GetAll() ([]models.TodoItem, error)
	GetById(itemId int) (models.TodoItem, error)
	Delete(itemId int) error
	Update(itemId int, input models.UpdateTodoItemInput) (models.TodoItem, error)
}
