package repositories

import (
	"golang-ture/ent"
	"golang-ture/internal/models"
)

type Repository struct {
	User
	TodoItem
}

type User interface {
	CreateUser(user models.SignInput) (int, error)
	GetUser(username, password string) (models.User, error)
	GetById(userId int) (user models.User, err error)
}

type TodoItem interface {
	Create(item models.TodoItem) (int, error)
	GetAll() ([]models.TodoItem, error)
	GetById(itemId int) (models.TodoItem, error)
	Delete(itemId int) error
	Update(itemId int, input models.UpdateTodoItemInput) (models.TodoItem, error)
}

func NewRepositoryStorage(storage *Storage) *Repository {
	storage.Data = make(map[string]map[string]string)
	return &Repository{
		TodoItem: NewTodoItemRepositoryStorage(storage),
	}
}

func NewRepository(DBClient *ent.Client) *Repository {
	return &Repository{
		User:     NewUserRepositoryEnt(DBClient),
		TodoItem: NewTodoItemRepositoryEnt(DBClient),
	}
}
