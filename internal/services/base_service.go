package services

import (
	"golang-ture/internal/models"
	"golang-ture/internal/repositories"
)

type Service struct {
	User
	TodoItem
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		User:     NewUserService(repos.User),
		TodoItem: NewTodoItemService(repos.TodoItem),
	}
}

type User interface {
	CreateUser(user models.SignInput) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	GetById(userId int) (user models.User, err error)
}

type TodoItem interface {
	Create(item models.TodoItem) (int, error)
	GetAll() ([]models.TodoItem, error)
	GetById(itemId int) (models.TodoItem, error)
	Delete(itemId int) error
	Update(itemId int, input models.UpdateTodoItemInput) (models.TodoItem, error)
}
