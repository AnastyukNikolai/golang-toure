package services

import (
	"golang-ture/internal/models"
	"golang-ture/internal/repositories"
)

type TodoItemService struct {
	repo repositories.TodoItem
}

func NewTodoItemService(repo repositories.TodoItem) *TodoItemService {
	return &TodoItemService{repo: repo}
}

func (s *TodoItemService) Create(item models.TodoItem) (string, error) {
	return s.repo.Create(item)
}

func (s *TodoItemService) GetAll() ([]models.TodoItem, error) {
	return s.repo.GetAll()
}

func (s *TodoItemService) GetById(itemId string) (models.TodoItem, error) {
	return s.repo.GetById(itemId)
}

func (s *TodoItemService) Delete(itemId string) error {
	return s.repo.Delete(itemId)
}

func (s *TodoItemService) Update(itemId string, input models.UpdateTodoItemInput) error {
	return s.repo.Update(itemId, input)
}
