package repositories

import (
	"encoding/json"
	"errors"
	"golang-ture/internal/models"
	_map "golang-ture/pkg/map"
)

type TodoItemRepository struct {
	storage *Storage
}

func NewTodoItemRepositoryStorage(storage *Storage) *TodoItemRepository {
	return &TodoItemRepository{storage: storage}
}

func (r *TodoItemRepository) Create(item models.TodoItem) (int, error) {
	guid := item.Id
	_, exist := r.storage.Data[TodoItemsKey][string(rune(guid))]
	if exist {
		return 0, errors.New("item already exist")
	}
	jsonItem, err := json.Marshal(item)
	if err != nil {
		return 0, err
	}
	newItem := make(map[string]string)
	newItem[string(rune(guid))] = string(jsonItem)
	r.storage.Lock()
	r.storage.Data[TodoItemsKey] = _map.MergeStringMaps(r.storage.Data[TodoItemsKey], newItem)
	r.storage.Unlock()

	return guid, nil
}

func (r *TodoItemRepository) GetAll() ([]models.TodoItem, error) {
	var items []models.TodoItem
	for _, value := range r.storage.Data[TodoItemsKey] {
		var todo models.TodoItem
		err := json.Unmarshal([]byte(value), &todo)
		if err != nil {
			return nil, err
		}
		items = append(items, todo)
	}
	return items, nil
}

func (r *TodoItemRepository) GetById(itemId int) (models.TodoItem, error) {
	var todo models.TodoItem
	item, ok := r.storage.Data[TodoItemsKey][string(rune(itemId))]
	if !ok {
		return todo, errors.New("item not found")
	}
	err := json.Unmarshal([]byte(item), &todo)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (r *TodoItemRepository) Delete(itemId int) error {
	_, ok := r.storage.Data[TodoItemsKey][string(rune(itemId))]
	if !ok {
		return errors.New("item not found")
	}
	r.storage.Lock()
	delete(r.storage.Data[TodoItemsKey], string(rune(itemId)))
	r.storage.Unlock()
	return nil
}

func (r *TodoItemRepository) Update(itemId int, input models.UpdateTodoItemInput) (models.TodoItem, error) {
	todo, err := r.GetById(itemId)
	if err != nil {
		return todo, err
	}

	if input.Title != nil {
		todo.ModifyTitle(*input.Title)
	}

	if input.Description != nil {
		todo.ModifyDescription(*input.Description)
	}

	if input.Status != nil {
		todo.ModifyStatus(*input.Status)
	}

	if input.Done != nil {
		todo.ModifyDone(*input.Done)

	}

	updatedTodo, err := json.Marshal(todo)
	if err != nil {
		return todo, err
	}

	r.storage.Lock()
	r.storage.Data[TodoItemsKey][string(rune(itemId))] = string(updatedTodo)
	r.storage.Unlock()

	return todo, nil
}
