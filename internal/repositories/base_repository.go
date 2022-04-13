package repositories

import (
	"golang-ture/internal/models"
)

type TodoItem interface {
	Create(item models.TodoItem) (string, error)
	GetAll() ([]models.TodoItem, error)
	GetById(itemId string) (models.TodoItem, error)
	Delete(itemId string) error
	Update(itemId string, input models.UpdateTodoItemInput) error
}

type Repository struct {
	TodoItem
}

func NewRepository(storage *Storage) *Repository {
	storage.Data = make(map[string]map[string]string)
	return &Repository{
		TodoItem: NewTodoItemRepository(storage),
	}
}
