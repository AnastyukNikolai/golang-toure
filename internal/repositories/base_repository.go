package repositories

import (
	"golang-ture/ent"
	"golang-ture/internal/models"
)

type TodoItem interface {
	Create(item models.TodoItem) (int, error)
	GetAll() ([]models.TodoItem, error)
	GetById(itemId int) (models.TodoItem, error)
	Delete(itemId int) error
	Update(itemId int, input models.UpdateTodoItemInput) (models.TodoItem, error)
}

type Repository struct {
	TodoItem
}

func NewRepositoryStorage(storage *Storage) *Repository {
	storage.Data = make(map[string]map[string]string)
	return &Repository{
		TodoItem: NewTodoItemRepositoryStorage(storage),
	}
}

func NewRepository(DBClient *ent.Client) *Repository {
	return &Repository{
		TodoItem: NewTodoItemRepositoryEnt(DBClient),
	}
}
