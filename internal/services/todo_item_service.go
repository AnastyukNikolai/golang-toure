package services

import (
	"golang-ture/internal/models"
	"golang-ture/internal/repositories"
)

type TodoItemService struct {
	repo repositories.TodoItem
}

// Create new instance of TodoItemService
func NewTodoItemService(repo repositories.TodoItem) *TodoItemService {
	return &TodoItemService{repo: repo}
}

// Create new TodoItem
func (s *TodoItemService) Create(item models.TodoItem) (int, error) {
	return s.repo.Create(item)
}

// Get all TodoItems from storage
func (s *TodoItemService) GetAll() ([]models.TodoItem, error) {
	return s.repo.GetAll()
}

// Get TodoItem from storage by Id
func (s *TodoItemService) GetById(itemId int) (models.TodoItem, error) {
	return s.repo.GetById(itemId)
}

// Delete TodoItem from storage by Id
func (s *TodoItemService) Delete(itemId int) error {
	return s.repo.Delete(itemId)
}

// Update TodoItem fields by Id
func (s *TodoItemService) Update(itemId int, input models.UpdateTodoItemInput) (models.TodoItem, error) {
	return s.repo.Update(itemId, input)
}
